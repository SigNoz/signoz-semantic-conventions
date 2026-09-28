---
paths:
- "docs/**"
- "templates/**"
---

# Docs

- `docs/` is generated from `model/` by `make weaver-docs`, using the templates in `templates/registry/markdown/`. Never edit `docs/` by hand: change the model or the templates and regenerate. CI fails when `docs/` doesn't match.
- The pages are for people using SigNoz telemetry. They carry no text about how they are built or where the YAML lives; notes for contributors go in HTML comments.
