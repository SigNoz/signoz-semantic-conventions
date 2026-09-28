---
paths:
- "docs/**"
- "templates/**"
---

# Docs

- `docs/` is generated from `model/` by `make weaver-docs`, using the templates in `templates/registry/markdown/`. Never edit `docs/` by hand: change the model or the templates and regenerate. CI fails when `docs/` doesn't match.
- Templates: `weaver.yaml` groups everything by the `model/` directory it came from; `area.md.j2` renders one page per area, `readme.md.j2` the index, `macros.j2` the tables and links. A new `model/<area>/` becomes a new page with no template change.
- The pages are for people using SigNoz telemetry. They carry no text about how they are built or where the YAML lives; notes for contributors go in HTML comments.
