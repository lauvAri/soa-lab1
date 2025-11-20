#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

if ! command -v swag >/dev/null 2>&1; then
  echo "swag command is not installed. Install it via 'go install github.com/swaggo/swag/cmd/swag@latest'." >&2
  exit 1
fi

# Ensure go commands invoked by swag work even when HOME is read-only.
if [[ -z "${GOCACHE:-}" ]]; then
  export GOCACHE="/tmp/go-cache"
fi
mkdir -p "${GOCACHE}"

# Avoid git safe.directory issues when go list runs under swag.
if [[ -z "${GOFLAGS:-}" ]]; then
  export GOFLAGS="-buildvcs=false"
elif [[ "${GOFLAGS}" != *"-buildvcs=false"* ]]; then
  export GOFLAGS="${GOFLAGS} -buildvcs=false"
fi

cd "${ROOT_DIR}"
swag init -g main.go -o docs "$@"
