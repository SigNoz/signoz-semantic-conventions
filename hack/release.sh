#!/usr/bin/env bash
# Release helpers for the registry, used by the release workflows (see RELEASE.md).
#
# Usage:
#   hack/release.sh version              Print the version in model/manifest.yaml.
#   hack/release.sh prepare <version>    Check <version> and set it in model/manifest.yaml.
#   hack/release.sh verify <tag>         Check <tag> is v<version> for the version in model/manifest.yaml.
#   hack/release.sh package <tag> <dir>  Package the registry for release <tag> into <dir>.
#   hack/release.sh bump <tag>           After release <tag>, set the next unreleased version; print it.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

readonly manifest="model/manifest.yaml"
readonly prefix="https://signoz.io/schemas/otel/"
# SemVer 2.0 without build metadata: no leading zeros, non-empty pre-release identifiers.
readonly semver='^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-(0|[1-9][0-9]*|[0-9]*[A-Za-z-][0-9A-Za-z-]*)(\.(0|[1-9][0-9]*|[0-9]*[A-Za-z-][0-9A-Za-z-]*))*)?$'

die() {
  echo "release: $*" >&2
  exit 1
}

current_version() {
  sed -nE "s#^schema_url: ${prefix}(.*)#\1#p" "${manifest}"
}

set_version() {
  sed -i.bak -E "s#^(schema_url: ${prefix}).*#\1$1#" "${manifest}"
  rm -f "${manifest}.bak"
}

check_version() {
  [[ "$1" =~ ${semver} ]] || die "'$1' is not a SemVer version (MAJOR.MINOR.PATCH with an optional -pre-release, no v prefix)"
}

check_tag() {
  [[ "$1" == v* ]] || die "tag '$1' has no v prefix"
  check_version "${1#v}"
}

case "${1:-}" in
  version)
    current_version
    ;;
  prepare)
    version="${2:?usage: hack/release.sh prepare <version>}"
    check_version "${version}"
    if git ls-remote --exit-code --tags origin "refs/tags/v${version}" >/dev/null 2>&1; then
      die "tag v${version} already exists"
    fi
    set_version "${version}"
    ;;
  verify)
    tag="${2:?usage: hack/release.sh verify <tag>}"
    check_tag "${tag}"
    [[ "${tag}" == "v$(current_version)" ]] || die "tag ${tag} doesn't match schema_url version $(current_version) in ${manifest}"
    ;;
  package)
    tag="${2:?usage: hack/release.sh package <tag> <dir>}"
    dir="${3:?usage: hack/release.sh package <tag> <dir>}"
    check_tag "${tag}"
    repo="${GITHUB_REPOSITORY:-SigNoz/signoz-semantic-conventions}"
    if ! out="$(weaver registry package --v2 --quiet -r model/ -o "${dir}" \
      --resolved-registry-uri "https://github.com/${repo}/releases/download/${tag}/resolved.yaml" 2>&1)"; then
      die "weaver registry package failed:"$'\n'"${out}"
    fi
    ;;
  bump)
    tag="${2:?usage: hack/release.sh bump <tag>}"
    check_tag "${tag}"
    version="${tag#v}"
    base="${version%%-*}"
    current="$(current_version)"
    if [[ "${current%%-*}" != "${base}" ]]; then
      # main already prepares another version, e.g. an older release was published.
      echo "release: main prepares ${current}, not ${base}; nothing to bump" >&2
      exit 0
    fi
    if [[ "${version}" == *-* ]]; then
      next="${base}-unreleased"
    else
      IFS=. read -r major minor _ <<<"${base}"
      next="${major}.$((minor + 1)).0-unreleased"
    fi
    set_version "${next}"
    echo "${next}"
    ;;
  *)
    sed -n '4,9s/^# \{0,1\}//p' "$0"
    exit 1
    ;;
esac
