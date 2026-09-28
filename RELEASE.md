# Releasing

## Versions

- The version is the last segment of `schema_url` in [`model/manifest.yaml`](model/manifest.yaml), with no `v` prefix: `https://signoz.io/schemas/otel/0.0.1-rc.1`. The OpenTelemetry schema URL spec requires a SemVer-ordered `MAJOR.MINOR.PATCH` there.
- The git tag and GitHub release are the version with a `v` prefix: `v0.0.1-rc.1`.
- A version with a suffix such as `-rc.1` is published as a pre-release.
- Between releases, `main` carries the next version with an `-unreleased` suffix.

## Setup

The three workflows authenticate as the primus GitHub App through the `PRIMUS_APP_ID` and `PRIMUS_PRIVATE_KEY` organization secrets, so their PRs and releases trigger other workflows. The app needs write access to contents and pull requests on this repository.

## Steps

1. Run the [prereleaser](.github/workflows/prereleaser.yaml) workflow with the version, e.g. `0.0.1-rc.1`. It opens a `chore(release): v<version>` PR that sets `schema_url`.
2. Review and merge that PR.
3. Tag the merge commit and push the tag:

   ```bash
   git switch main && git pull
   git tag v0.0.1-rc.1
   git push origin v0.0.1-rc.1
   ```

4. The [releaser](.github/workflows/releaser.yaml) workflow runs on the tag. It checks that the tag matches `schema_url`, runs `make yaml-lint` and `make weaver-check`, packages the registry with `weaver registry package`, and publishes the GitHub release with `manifest.yaml` and `resolved.yaml` attached and notes generated from the merged PRs.
5. The [postreleaser](.github/workflows/postreleaser.yaml) workflow opens a PR that sets `schema_url` back to an unreleased version: `0.0.1-unreleased` after `v0.0.1-rc.1`, `0.2.0-unreleased` after `v0.1.0`. Review and merge it.
