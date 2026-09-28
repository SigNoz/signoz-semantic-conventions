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
- Test a change by hand: build a scratch `model/` that breaks the rule and run weaver from its parent directory (`-r .` loads nothing), then run `make weaver-check` on the real model.
