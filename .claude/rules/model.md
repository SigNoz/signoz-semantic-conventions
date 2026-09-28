---
paths:
- "model/**"
---

# Model files

## YAML

- Weaver's v1 `groups:` format only; never `file_format: definition/2`.
- `brief:` is always a `>` block and `note:` a `|` block. One blank line between groups.
- Attributes are defined only in `registry.*` attribute groups; every other group references them with `- ref:`.
- A `ref` never sets `stability` or `deprecated`; weaver silently applies them to that signal. It may refine only `brief`, `note`, `examples`, `tag`, `annotations`, `requirement_level`, `sampling_relevant` and (entities) `role`. An inline `ref` beats one inherited through `extends`; along an `extends` chain the nearest wins.
- Events have no `body:`; describe any display-message body in the `note`.
- Examples: `string[]` takes a list of lists (`- - "a"`); template examples are values, not `key=value`.
- `imports:` match spans and attribute groups by group id (`span.http.client`, `client`), metrics and events by name, entities by type.

## Names

- SigNoz names start with `signoz.`. An upstream key is referenced, never redefined, and nothing new goes under an upstream namespace such as `http.` or `db.`.
- Two attributes, two metrics or two events never share a name.

## Changing what exists

- Never delete or rename an emitted name in place: move the old definition to `model/<area>/deprecated/` with a structured `deprecated:` block.
- A stable convention never changes an attribute's type, an enum member's value (and never loses a member), a metric's unit or instrument, a span's kind or an entity's identity. Stable attributes are never dropped from stable signals, and a stable metric only gains `opt_in` attributes.
- A rename that only swaps `_` and `.` (`rule_id` to `rule.id`) gives both names one code constant; set `annotations.code_generation.exclude: true` on the old one.

## Exceptions

- Only for names already emitted; anything new gets a name that passes.
- List the rule id under `annotations.signoz.policy_exceptions` for our rules, `annotations.stability.policy_exceptions` (finding id minus `stability_`) or `annotations.naming_conventions.policy_exceptions` (only `metric_namespace_collision`) for upstream's.

## Briefs and notes

- A `brief` is one noun phrase, usually starting with "The".
- Say what a thing is, never how it is produced, stored or queried, or why it was designed that way.
- Group notes are one or two lines. Don't restate the YAML: requirement level, `extends`, event name or type.
- A possibly sensitive value gets a `> [!WARNING]` note saying what is sensitive, with no operating advice.

## Before committing

Run `make yaml-lint weaver-check weaver-docs` and commit the regenerated `docs/`.
