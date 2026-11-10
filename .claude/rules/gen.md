---
paths:
- "gen/**"
- "templates/registry/go/**"
- "templates/registry/python/**"
- "templates/registry/javascript/**"
---

# Generated code

- `gen/` is generated from `model/` by `make weaver-gen` (`weaver-go`, `weaver-python`, `weaver-javascript`), using the templates in `templates/registry/<language>/`. Never edit `gen/` by hand: change the model or the templates and regenerate. CI fails when `gen/` doesn't match. Only `gen/go/go.mod`, `gen/go/go.sum` and `gen/python/pyproject.toml` are hand-written.
- The layout is `gen/<language>/sigsemconv/`, and the package is named `sigsemconv` in every language. Each language's manifest sits in `gen/<language>/`.
- Every language groups definitions by the `model/` directory they came from. A new `model/<area>/` becomes new files with no template change.
- The templates follow the OpenTelemetry templates for each language, linked at the top of each `weaver.yaml`. Keep them close to upstream, and list every difference in that header comment.
- Go names write the words under `acronyms` in `templates/registry/go/weaver.yaml` in a fixed casing (`ID`, `SigNoz`). The list holds only what the registry's names use: when a new name contains an acronym that isn't listed, add it.
- The version in the generated code comes from `schema_url` in `model/manifest.yaml`, so a version change regenerates `gen/`.
- Nothing is published to a package registry. Consumers install from the release tag; Go resolves the module in `gen/go` from the tag `gen/go/v<version>`.
