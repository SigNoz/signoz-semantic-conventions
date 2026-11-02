---
paths:
- "semconv/go/**"
- "templates/registry/go/**"
---

# Go module

- `semconv/go/` is generated from `model/` by `make weaver-go`, using the templates in `templates/registry/go/`. Never edit it by hand: change the model or the templates and regenerate. CI fails when it doesn't match.
- Keep `go 1.23` in `go.mod.j2` and import nothing beyond the standard library, so projects on older Go can import the module.
- Release `v<version>` tags the module as `semconv/go/v<version>`.
