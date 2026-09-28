<p align="center">
  <a href="https://signoz.io" target="_blank">
    <img alt="SigNoz" src="https://signoz.io/img/SigNozLogo-orange.svg" width="120">
  </a>
</p>

<h1 align="center" style="border-bottom: none">SigNoz Semantic Conventions</h1>

<p align="center">
  <a href="https://github.com/SigNoz/signoz-semantic-conventions/releases"><img alt="GitHub Release" src="https://img.shields.io/github/v/release/SigNoz/signoz-semantic-conventions?include_prereleases"></a>
  <a href="LICENSE"><img alt="License: AGPL v3" src="https://img.shields.io/badge/License-AGPL%20v3-blue.svg"></a>
  <a href="https://signoz.io/slack"><img alt="Slack" src="https://img.shields.io/badge/Slack-SigNoz-4A154B?logo=slack"></a>
</p>

<p align="center">The names, types and meanings of the telemetry SigNoz emits and reads.</p>

## Overview

The conventions are an [OpenTelemetry Weaver](https://github.com/open-telemetry/weaver) registry built on the [OpenTelemetry semantic conventions](https://github.com/open-telemetry/semantic-conventions). The YAML definitions are in [`model/`](model), one folder per area, and the [docs](docs/README.md) are generated from them.

| Area | Covers |
| --- | --- |
| [`audit`](docs/audit.md) | Audit log events for operations on SigNoz resources |
| [`genai`](docs/genai.md) | The cost of generative AI model calls |
| [`identity`](docs/identity.md) | Organizations and principals |
| [`meter`](docs/meter.md) | Usage metrics for ingested telemetry |
| [`resource`](docs/resource.md) | SigNoz resources such as dashboards, alert rules and ingestion keys |
| [`spanmetrics`](docs/spanmetrics.md) | Metrics derived from spans for APM |
| [`workspace`](docs/workspace.md) | Workspaces and their ingestion keys |

## Using the conventions

To build on these conventions in your own Weaver registry, depend on a [release](https://github.com/SigNoz/signoz-semantic-conventions/releases) in your `manifest.yaml`, then reference the attributes by key:

```yaml
dependencies:
- schema_url: https://signoz.io/schemas/otel/<version>
  registry_path: https://github.com/SigNoz/signoz-semantic-conventions/archive/refs/tags/v<version>.zip[model]
```

Each release also carries the resolved registry (`resolved.yaml`) for tools that read it directly. [RELEASE.md](RELEASE.md) describes how versions and releases work.

## Contributing

Open an [issue](https://github.com/SigNoz/signoz-semantic-conventions/issues) or a pull request, ask in `#contributing` on the [SigNoz Slack](https://signoz.io/slack), or share ideas in [GitHub Discussions](https://github.com/SigNoz/signoz/discussions).

Changes are checked with [weaver](https://github.com/open-telemetry/weaver/releases) and [yamllint](https://github.com/adrienverge/yamllint): run `make yaml-lint`, `make weaver-check` and `make weaver-docs`, and commit the regenerated `docs/` with your change. CI runs the same checks.

## License

Licensed under the [GNU Affero General Public License v3.0](LICENSE).
