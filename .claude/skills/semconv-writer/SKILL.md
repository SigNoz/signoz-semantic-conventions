---
name: semconv-writer
description: Write, review and evolve OpenTelemetry-style semantic conventions for the SigNoz registry (the weaver-validated YAML under model/ in this repo). Use it whenever the user wants to add or change attributes, spans, metrics, events, entities/resources, enum values, deprecations or any registry YAML; asks how to name a metric or attribute, which requirement level, unit, instrument or stability to use, whether something already exists upstream in OTel, or how to run weaver/yamllint checks or the rego policies; or mentions semconv, semantic conventions, OTel conventions, registry.yaml, weaver, telemetry schema or instrumentation contracts - even when the words "semantic convention" never appear.
---

# SigNoz semantic-convention writer

## Repo

- Registry `signoz`, manifest `model/manifest.yaml`; its `dependencies` entry pins the upstream OTel semconv release.
- Hard constraints are in `.claude/rules/` (`model.md`, `docs.md`, `policies.md`, `pull-requests.md`).
- Make targets: `make yaml-lint`, `make weaver-check`, `make weaver-docs`; CI runs all three with the weaver version pinned in `.github/workflows/ci.yaml`.

## SigNoz decisions

- One folder per functional area under `model/<area>/`: `registry.yaml` (definitions, `registry.signoz.<area>`), then `spans.yaml`, `metrics.yaml`, `events.yaml`, `entities.yaml`, `deprecated/`. Each area is one `docs/` page.
- Shared refinements go in an `attributes.signoz.<area>.common` group that signals `extends:`.
- Live areas (meter, spanmetrics) are recorded exactly as emitted, including old upstream names (`deployment.environment`, `db.system`, `db.name`, `http.status_code`); rule breaks are exempted per definition, never renamed.
- Every span type gets a `{operation}.duration` histogram in `s`, with `error.type` on failure only.
- A new attribute needs a clear end-user benefit, a signal that carries it, and a way for instrumentation to get it.
- A `note` adds only what the brief can't; omit it otherwise. Group briefs stay generic.

## Workflow

1. Frame the use case and pick the signal, broad signals first (upstream: t-shaped signals).
2. Reuse before defining: `grep -rn "id: " model/`, then `python3 .claude/skills/semconv-writer/scripts/upstream.py attrs '<regex>'` (also `show <key>`, `metrics`, `events`, `entities`, `spans`, `namespaces`), then upstream `main` and open PRs.
3. Name, following upstream naming and `.claude/rules/model.md`.
4. Write the YAML, copying the shape of a file in `model/`.
5. Validate per `.claude/rules/model.md`; ignore the `definition/2` "is not yet stable" warnings, which come from the upstream dependency.
6. Open the PR per `.claude/rules/pull-requests.md`; list the conventions, link the use case, name the reused upstream keys.

## Upstream docs

Fetch at the pinned version: `gh api "repos/open-telemetry/semantic-conventions/contents/<path>?ref=v<version>" --jq .content | base64 -d`, version from `model/manifest.yaml`. Weaver docs: `repos/open-telemetry/weaver`, version from `.github/workflows/ci.yaml`.

| Topic | Path |
| --- | --- |
| Attributes, enums, spans, prototyping | `docs/how-to-write-conventions/README.md` |
| Picking signals | `docs/how-to-write-conventions/t-shaped-signals.md` |
| Entities | `docs/how-to-write-conventions/resource-and-entities.md` |
| Status metrics | `docs/how-to-write-conventions/status-metrics.md` |
| Naming | `docs/general/naming.md` |
| Requirement levels | `docs/general/attribute-requirement-level.md`, `docs/general/signal-requirement-level.md` |
| Metrics, units, instruments | `docs/general/metrics.md` |
| Spans | `docs/general/trace.md` |
| Events | `docs/general/events.md` |
| Errors, `error.type` | `docs/general/recording-errors.md` |
| Exceptions | `docs/exceptions/exceptions-logs.md`, `docs/exceptions/exceptions-spans.md` |
| Stability, groups | `docs/general/semantic-convention-groups.md` |
| YAML syntax (weaver) | `schemas/semconv-syntax.md` |
| Deprecation (weaver) | `docs/schema-changes.md` |
| Upstream policies | `open-telemetry/opentelemetry-weaver-packages`, `policies/check/*/README.md` |
