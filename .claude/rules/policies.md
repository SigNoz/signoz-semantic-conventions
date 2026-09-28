---
paths:
- "policies/**"
---

# Rego policies

- Every finding fails `make weaver-check`, whatever its `level`, so only write rules worth failing on.
- `before_resolution/` rules see one parsed model file at a time (weaver runs them without `--v2`); `after_resolution/` rules see the resolved registry (`--v2`). Put a rule where its input is.
- Name helpers `signoz_*`: our files share one package with the upstream policy files loaded next to them.
- A rule that shipped names may need to break checks `not signoz_excepted(obj, "<rule id>")`.
- `after_resolution/names.rego` is a copy of upstream's `names.rego` at the commit pinned in the `Makefile`. When that commit changes, copy it again and re-apply the exception checks.
- Test a change by hand: in a scratch directory, create `model/` with a copy of `model/manifest.yaml` and one file that breaks the rule, then from that directory run `weaver registry check -r model/ --policy <repo>/policies/before_resolution` (or `weaver registry check -r model/ --v2 --policy <repo>/policies/after_resolution` for a resolved-registry rule) and confirm the finding appears. Always pass `-r model/`: `-r .` loads no files. Then run `make weaver-check` on the real model.
