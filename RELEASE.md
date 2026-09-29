# Releasing

## Versions

- The version is the last segment of `schema_url` in [`model/manifest.yaml`](model/manifest.yaml), with no `v` prefix: `https://signoz.io/schemas/otel/0.0.1-rc.1`. The OpenTelemetry schema URL spec requires a SemVer-ordered `MAJOR.MINOR.PATCH` there.
- The git tag and GitHub release are the version with a `v` prefix: `v0.0.1-rc.1`.
- A version with a suffix such as `-rc.1` is a pre-release.
- Between releases, `main` carries the next version with an `-unreleased` suffix.

## Steps

The workflows run [`hack/release.sh`](hack/release.sh) for the version checks and edits; run it with no arguments for usage.

1. Run the [prereleaser](.github/workflows/prereleaser.yaml) workflow with the version, e.g. `0.0.1-rc.1`. It opens a `chore(release): v<version>` PR that sets `schema_url`.
2. Review and merge that PR.
3. [Create the release](https://github.com/SigNoz/signoz-semantic-conventions/releases/new) in the GitHub UI: tag `v<version>` on the merge commit, title `v<version>`, generated release notes, and "Set as a pre-release" when the version has a suffix. Publish it.
4. The [releaser](.github/workflows/releaser.yaml) workflow runs on the published release. It checks that the tag matches `schema_url`, packages the registry with `weaver registry package`, and attaches `manifest.yaml` and `resolved.yaml` to the release. To attach them again, run it by hand with the tag.
5. The [postreleaser](.github/workflows/postreleaser.yaml) workflow opens a PR that sets `schema_url` back to an unreleased version: `0.0.1-unreleased` after `v0.0.1-rc.1`, `0.2.0-unreleased` after `v0.1.0`. Review and merge it.
