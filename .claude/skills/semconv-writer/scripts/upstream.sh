#!/usr/bin/env bash
# Look things up in the upstream semconv release pinned in model/manifest.yaml.
# Usage: upstream.sh attrs|metrics|events|entities|spans <regex>
#        upstream.sh show <attribute key>
#        upstream.sh refresh
# Needs weaver, yq and jq.
set -euo pipefail

root="$(git rev-parse --show-toplevel)"
registry="$(yq '.dependencies[0].registry_path' "$root/model/manifest.yaml")"
version="$(yq '.dependencies[0].schema_url' "$root/model/manifest.yaml" | sed 's#.*/##')"
json="${TMPDIR:-/tmp}/semconv-upstream-$version.json"

if [[ "${1:-}" == refresh ]]; then rm -f "$json"; fi
if [[ ! -f "$json" ]]; then
  # Build in a temp dir, then rename the file into place, so a concurrent lookup never reads it half-written.
  tmp="$(mktemp -d "${TMPDIR:-/tmp}/semconv-upstream.XXXXXX")"
  trap 'rm -rf "$tmp"' EXIT
  if ! out="$(weaver registry package --v2 --quiet --resolved-registry-uri unused -o "$tmp" -r "$registry" 2>&1)"; then
    echo "$out" >&2; exit 1
  fi
  yq -o=json . "$tmp/resolved.yaml" >"$tmp/resolved.json"
  mv -f "$tmp/resolved.json" "$json"
fi

# The catalog repeats an attribute once per refinement; registry.attributes indexes the canonical one.
# shellcheck disable=SC2016 # jq variables, not shell ones
attrs='.attribute_catalog as $c | .registry.attributes[] | $c[.]'
pattern="${2:-.}"

case "${1:-}" in
  attrs)    jq -r --arg p "$pattern" "$attrs"' | select(.key | test($p)) | "\(.key)\t\(.stability)\(if .deprecated then " DEPRECATED" else "" end)\t\(.brief | gsub("\n"; " "))"' "$json" ;;
  show)     jq --arg k "$pattern" "$attrs"' | select(.key == $k)' "$json" ;;
  metrics)  jq -r --arg p "$pattern" '.registry.metrics[] | select(.name | test($p)) | "\(.name)\t\(.instrument)\t\(.unit)\t\(.stability)"' "$json" ;;
  events)   jq -r --arg p "$pattern" '.registry.events[] | select(.name | test($p)) | "\(.name)\t\(.stability)"' "$json" ;;
  entities) jq -r --arg p "$pattern" '.registry.entities[] | select(.type | test($p)) | "\(.type)\t\(.stability)"' "$json" ;;
  spans)    jq -r --arg p "$pattern" '.registry.spans[] | select(.type | test($p)) | "\(.type)\t\(.kind)\t\(.stability)"' "$json" ;;
  refresh)  echo "cached $json" ;;
  *)        sed -n '2,6s/^# //p' "$0"; exit 1 ;;
esac
