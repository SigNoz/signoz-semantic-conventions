---
paths:
- "policies/**"
---

# Rego policies

## How `make weaver-check` runs them

- First a grep fails on a group `id` or `metric_name` defined twice in `model/` (weaver only warns and keeps one; rego can't see it). Then two weaver runs. First, without `--v2`, `before_resolution/` rules run once per model file on weaver's parsed form (`input.groups`, `input.file_format`). Then, with `--v2`, `after_resolution/` rules and the upstream `opentelemetry-weaver-packages` files pinned in the `Makefile` see the resolved registry: `input.registry` (ours, refs resolved, `extends` merged) and `input.dependencies` (every upstream definition, keyed by schema URL).
- `before_resolution` rules are skipped under `--v2`. Rules about the file as written (group ids and types, inline attribute definitions, missing stability, and anything resolution loses, below) go in `before_resolution/`; everything else in `after_resolution/`.
- Every finding fails the check, whatever its `level`, so only write rules worth failing on.
- `after_resolution/names.rego` replaces upstream's; its header lists our two changes (`signoz_excepted` checks, and `naming_convention_attr_name` checking the full name format). When the commit pinned in the `Makefile` changes, copy it again and re-apply both.

## Writing a rule

- Name helpers `signoz_*`: `after_resolution` shares its package with the upstream files loaded next to it.
- Build findings with `signoz_finding` (`before_resolution`) or `signoz_attribute_finding` / `signoz_signal_finding` (`after_resolution`) from `lib.rego`.
- A rule that live names may break checks `not signoz_excepted(obj, "<rule id>")`. The annotation reaches the resolved registry on the definition and on every signal that refs it; upstream rules never read it.
- Resolved-run findings carry `Provenance: model/`, not the file; name the definition in the message.
- Don't duplicate what weaver already fails on: a `ref` that sets `type`, a bare `conditionally_required` without text, a duplicate attribute id.
- Weaver only warns (exit 0) on missing stability, `prefix`, a missing group `type`, a string attribute without examples, `deprecated: "text"` (it arrives as reason `unspecified`), flat `string[]` examples, a span without `span_kind`, `key=value` template examples and an import that matches nothing. Our rules fail on all but the last two, and on `stability`/`deprecated` set on a `ref`, which weaver silently applies.

## Weaver input quirks

- Weaver rewrites input before any policy sees it, so rules for these can't exist: `stability: experimental` arrives as `development`, `type: resource` as `entity`, and a deprecation without `note` gets `note: "Obsoleted."`.
- Resolution drops event bodies and `sampling_relevant` on non-span attributes, and turns a missing entity `role` into descriptive (an entity whose refs all lack one gets an empty identity); check those in `before_resolution/`.
- `before_resolution` input has defaults filled in (`requirement_level: "recommended"`, `note: ""`, `annotations: null`), so an explicit value and an omitted one look the same.
- Shared `attributes.*` groups are absent from `input.registry.attribute_groups`; their attributes appear only on the signals that pull them in.
- Span `events:` lists aren't validated; a nonexistent id passes.
- Regorus differs from OPA: `some x in <string>` is an error, not "no match". Guard fields that can be a string or an object (`requirement_level`, `type`) with `is_object`.
- `--future` turns upstream's `definition/2` warnings into errors, so it can't be used with the current dependency.
- Weaver has no built-in compatibility check; `--baseline-registry` only feeds comparison policies (upstream's `backwards-compatibility` package, once we publish releases).

## Testing

- Test by hand: in a scratch dir, create `model/` with a copy of `model/manifest.yaml` and one file that breaks the rule. Run `weaver registry check -r model/ --policy <repo>/policies/before_resolution` (add `--v2` and use `after_resolution` for a resolved rule), adding `--diagnostic-format json --diagnostic-stdout` to read findings as JSON. Always pass `-r model/`: `-r .` loads no files. Then run `make weaver-check` on the real model.
- To see what a rule receives, add a temporary `deny` rule whose message is `json.marshal(<value>)`.
