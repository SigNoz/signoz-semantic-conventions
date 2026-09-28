#!/usr/bin/env python3
"""
upstream.py - look things up in the upstream OpenTelemetry registry that
model/manifest.yaml pins as a dependency.

Reuse beats invention: before defining an attribute, metric, event or entity,
check whether the pinned upstream version already has it (and under which
name - old versions use names that were later renamed).

Usage (run from anywhere inside the repo):
  upstream.py attrs   <regex>      # attributes whose key matches
  upstream.py show    <key>        # full definition of one attribute
  upstream.py metrics <regex>
  upstream.py events  <regex>
  upstream.py entities <regex>     # names usable in entity_associations
  upstream.py spans   <regex>
  upstream.py namespaces           # root namespaces upstream already owns
  upstream.py info                 # which upstream version is pinned, cache path
  upstream.py refresh              # re-resolve the dependency into the cache

The dependency is resolved once with `weaver registry resolve` and cached under
$XDG_CACHE_HOME/semconv-writer (override with $SEMCONV_WRITER_CACHE).
"""
from __future__ import annotations

import argparse
import hashlib
import json
import os
import re
import shutil
import subprocess
import sys

sys.dont_write_bytecode = True  # keep scripts/ free of __pycache__


# --------------------------------------------------------------------------- #
# bootstrap helpers
# --------------------------------------------------------------------------- #
def ensure_yaml():
    """Import PyYAML; if missing, re-exec under yamllint's interpreter, which
    always ships PyYAML (the repo already requires yamllint)."""
    try:
        import yaml  # type: ignore

        return yaml
    except ImportError:
        pass
    if os.environ.get("SEMCONV_WRITER_REEXEC") != "1":
        yl = shutil.which("yamllint")
        if yl:
            with open(yl, "r", encoding="utf-8", errors="ignore") as fh:
                first = fh.readline().strip()
            if first.startswith("#!"):
                argv = first[2:].split() + [sys.argv[0]] + sys.argv[1:]
                os.environ["SEMCONV_WRITER_REEXEC"] = "1"
                try:
                    os.execvp(argv[0], argv)
                except OSError:
                    pass
    sys.exit(
        "PyYAML is required. Install it (python3 -m pip install pyyaml) or "
        "install yamllint (brew install yamllint) so its interpreter can be reused."
    )


def find_repo_root(start: str | None = None) -> str:
    p = os.path.abspath(start or os.getcwd())
    while True:
        if os.path.isfile(os.path.join(p, "model", "manifest.yaml")):
            return p
        parent = os.path.dirname(p)
        if parent == p:
            break
        p = parent
    here = os.path.dirname(os.path.abspath(__file__))
    cand = os.path.abspath(os.path.join(here, "..", "..", "..", ".."))
    if os.path.isfile(os.path.join(cand, "model", "manifest.yaml")):
        return cand
    sys.exit("cannot find model/manifest.yaml - run from inside the registry repo")


def read_manifest(root: str) -> dict:
    yaml = ensure_yaml()
    with open(os.path.join(root, "model", "manifest.yaml"), encoding="utf-8") as fh:
        return yaml.safe_load(fh) or {}


def cache_dir() -> str:
    d = os.environ.get("SEMCONV_WRITER_CACHE")
    if not d:
        base = os.environ.get("XDG_CACHE_HOME") or os.path.expanduser("~/.cache")
        d = os.path.join(base, "semconv-writer")
    os.makedirs(d, exist_ok=True)
    return d


def resolve_dependency(dep: dict, refresh: bool = False) -> str:
    """Resolve one manifest dependency to JSON with weaver; return cache path."""
    schema_url = dep.get("schema_url", "")
    registry_path = dep.get("registry_path", "")
    if not registry_path:
        sys.exit(f"dependency {schema_url} has no registry_path in the manifest")
    key = hashlib.sha1(f"{schema_url}|{registry_path}".encode()).hexdigest()[:12]
    out = os.path.join(cache_dir(), f"upstream-{key}.json")
    if os.path.isfile(out) and os.path.getsize(out) > 0 and not refresh:
        return out
    if not shutil.which("weaver"):
        sys.exit("weaver is not installed (see https://github.com/open-telemetry/weaver)")
    cmd = ["weaver", "registry", "resolve", "-r", registry_path, "-f", "json", "-o", out]
    proc = subprocess.run(cmd, stdout=subprocess.DEVNULL, stderr=subprocess.PIPE, text=True)
    if proc.returncode != 0 or not os.path.isfile(out):
        sys.exit(f"weaver failed to resolve {registry_path}:\n{proc.stderr[-2000:]}")
    return out


def _norm_stability(s):
    return {"experimental": "development"}.get(s, s)


def build_index(resolved_path: str, schema_url: str) -> dict:
    with open(resolved_path, encoding="utf-8") as fh:
        data = json.load(fh)
    groups = data.get("groups") or (data.get("registry") or {}).get("groups") or []
    attrs: dict[str, dict] = {}
    metrics: dict[str, dict] = {}
    events: dict[str, dict] = {}
    entities: dict[str, dict] = {}
    spans: dict[str, dict] = {}
    # registry groups first so the canonical definition wins over refinements
    ordered = sorted(groups, key=lambda g: 0 if str(g.get("id", "")).startswith("registry.") else 1)
    for g in ordered:
        gid = g.get("id", "")
        gtype = g.get("type")
        for a in g.get("attributes") or []:
            name = a.get("name")
            if not name:
                continue
            rec = dict(a)
            rec["stability"] = _norm_stability(rec.get("stability"))
            rec["defined_in"] = gid
            if name not in attrs:
                attrs[name] = rec
        if gtype == "metric" and g.get("metric_name"):
            metrics[g["metric_name"]] = g
        elif gtype == "event":
            events[g.get("name") or gid] = g
        elif gtype in ("entity", "resource"):
            entities[g.get("name") or gid] = g
        elif gtype == "span":
            spans[gid] = g
    namespaces = sorted({k.split(".")[0] for k in attrs} | {m.split(".")[0] for m in metrics})
    return {
        "schema_url": schema_url,
        "resolved_path": resolved_path,
        "attrs": attrs,
        "metrics": metrics,
        "events": events,
        "entities": entities,
        "spans": spans,
        "namespaces": namespaces,
    }


def load_upstream_index(root: str | None = None, refresh: bool = False) -> dict | None:
    """Index of everything the pinned upstream dependency defines, or None if
    the manifest has no dependencies."""
    root = root or find_repo_root()
    manifest = read_manifest(root)
    deps = manifest.get("dependencies") or []
    if not deps:
        return None
    merged = None
    for dep in deps:
        idx = build_index(resolve_dependency(dep, refresh), dep.get("schema_url", ""))
        if merged is None:
            merged = idx
        else:
            for k in ("attrs", "metrics", "events", "entities", "spans"):
                for name, rec in idx[k].items():
                    merged[k].setdefault(name, rec)
            merged["namespaces"] = sorted(set(merged["namespaces"]) | set(idx["namespaces"]))
    return merged


# --------------------------------------------------------------------------- #
# CLI
# --------------------------------------------------------------------------- #
def _type_str(t):
    if isinstance(t, dict) and "members" in t:
        vals = [str(m.get("value")) for m in t["members"][:6]]
        more = "" if len(t["members"]) <= 6 else ", ..."
        return f"enum[{', '.join(vals)}{more}]"
    return str(t)


def _one_line(s, n=90):
    s = " ".join(str(s or "").split())
    return s if len(s) <= n else s[: n - 3] + "..."


def cmd_attrs(idx, pattern):
    rx = re.compile(pattern)
    rows = [(k, v) for k, v in sorted(idx["attrs"].items()) if rx.search(k)]
    if not rows:
        print(f"no upstream attribute matches /{pattern}/")
        return 1
    for k, v in rows:
        dep = " DEPRECATED" if v.get("deprecated") else ""
        print(f"{k:<45} {_type_str(v.get('type')):<22} {v.get('stability') or '?':<12}{dep} {_one_line(v.get('brief'))}")
    return 0


def cmd_show(idx, key):
    v = idx["attrs"].get(key)
    if not v:
        print(f"{key} is not defined upstream ({idx['schema_url']})")
        near = [k for k in idx["attrs"] if key.split(".")[0] == k.split(".")[0]]
        if near:
            print("attributes in the same root namespace:", ", ".join(sorted(near)[:25]))
        return 1
    print(json.dumps({k: v[k] for k in v if k != "lineage"}, indent=2, default=str))
    return 0


def _cmd_named(idx, bucket, pattern, label):
    rx = re.compile(pattern)
    rows = [(k, v) for k, v in sorted(idx[bucket].items()) if rx.search(k)]
    if not rows:
        print(f"no upstream {label} matches /{pattern}/")
        return 1
    for k, v in rows:
        extra = ""
        if bucket == "metrics":
            extra = f"{v.get('instrument', '?'):<14} {v.get('unit', '?'):<12}"
        print(f"{k:<50} {extra}{_norm_stability(v.get('stability')) or '-':<12} {_one_line(v.get('brief'))}")
    return 0


def main(argv=None) -> int:
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("command", choices=["attrs", "show", "metrics", "events", "entities", "spans", "namespaces", "info", "refresh"])
    ap.add_argument("pattern", nargs="?", default=".")
    ap.add_argument("--root", help="repo root (default: nearest ancestor with model/manifest.yaml)")
    args = ap.parse_args(argv)

    root = find_repo_root(args.root)
    idx = load_upstream_index(root, refresh=(args.command == "refresh"))
    if idx is None:
        print("manifest declares no dependencies; nothing to look up")
        return 0
    if args.command in ("info", "refresh"):
        print(f"repo root:       {root}")
        print(f"upstream schema: {idx['schema_url']}")
        print(f"resolved cache:  {idx['resolved_path']}")
        print(f"attributes: {len(idx['attrs'])}  metrics: {len(idx['metrics'])}  events: {len(idx['events'])}  entities: {len(idx['entities'])}  spans: {len(idx['spans'])}")
        return 0
    if args.command == "namespaces":
        print("\n".join(idx["namespaces"]))
        return 0
    if args.command == "attrs":
        return cmd_attrs(idx, args.pattern)
    if args.command == "show":
        return cmd_show(idx, args.pattern)
    return _cmd_named(idx, args.command, args.pattern, args.command)


if __name__ == "__main__":
    sys.exit(main())
