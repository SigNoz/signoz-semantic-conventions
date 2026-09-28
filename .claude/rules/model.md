---
paths:
- "model/**"
---

# Model files

These hold for every change under `model/`.

## YAML

- Weaver's v1 `groups:` format only; never `file_format: definition/2`.
- Block lists, never `[a, b]`. Sequences are not indented under their parent key. `brief:` is always a `>` block and `note:` always a `|` block. One blank line between groups.
- Attributes are defined only in `registry.*` attribute groups; every other group references them with `- ref:`.
- Every group has an explicit `type`. Every signal, attribute definition and enum member has a `stability`; new things start at `development`.
- Every entity attribute sets `role: identifying` or `role: descriptive`. An omitted role counts as descriptive, and an entity with no identifying attribute fails.
- A `ref` never sets `type`, `stability` or `deprecated`; weaver silently applies the last two to that signal.
- No `prefix:`. No `body:` on events; describe any display-message body in the event's `note`.

## Names

- Names match `^[a-z][a-z0-9]*([._][a-z0-9]+)*$`.
- SigNoz-owned names start with `signoz.`.
- Never redefine an upstream key (reference it instead), and never define anything under an upstream namespace such as `http.` or `db.`.
- Two attributes, two metrics or two events never share a name.

## Changing what exists

- The registry is append-only. Never delete or rename in place: move the old definition to `model/<area>/deprecated/` with a structured `deprecated:` block.
- A stable convention never changes an attribute's type, an enum member's value (and never loses a member), a metric's unit or instrument, a span's kind or an entity's identity. Stable attributes are never dropped from stable signals, and a stable metric only gains `opt_in` attributes.
- `annotations.signoz.policy_exceptions` is only for names already emitted. Anything new gets a name that passes.

## Briefs and notes

- A `brief` is one noun phrase, usually starting with "The".
- Say what a thing is. Never how it is produced, stored or queried, and never why it was designed that way; that belongs in the PR description.
- Group notes are one or two lines.
- Don't restate what the YAML already says: requirement level, `extends`, event name or type.
- A value that may be sensitive gets a `> [!WARNING]` block in its note saying what is sensitive, with no operating advice.

## Before committing

Run `make yaml-lint`, `make weaver-check` and `make weaver-docs`, and commit the regenerated `docs/` with the change.
