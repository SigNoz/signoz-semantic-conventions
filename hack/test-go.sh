#!/usr/bin/env bash
# Checks that the generated Go module works for a consumer.
#
# Vets and tests gen/go, then builds a throwaway module that imports every
# package of it the way a consumer would.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

readonly module="${PWD}/gen/go"

cd "${module}"
go vet ./...
go test ./...

path="$(go list -m)"
# Every package a consumer can import; internal ones are not importable.
packages="$(go list ./... | grep -v '/internal/')"

dir="$(mktemp -d)"
trap 'rm -rf "${dir}"' EXIT
cd "${dir}"

{
  echo "package main"
  echo
  echo "import ("
  for package in ${packages}; do
    echo "	_ \"${package}\""
  done
  echo ")"
  echo
  echo "func main() {}"
} >main.go

go mod init consumer >/dev/null 2>&1
go mod edit -require "${path}@v0.0.0" -replace "${path}=${module}"
go mod tidy >/dev/null 2>&1
go build ./...
echo "checked $(echo "${packages}" | wc -l | tr -d ' ') packages"
