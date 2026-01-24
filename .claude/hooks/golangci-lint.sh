#!/usr/bin/env bash
set -euo pipefail

# Skip if not a Go module yet
[[ -f "go.mod" ]] || exit 0

command -v golangci-lint >/dev/null || exit 0
golangci-lint run ./...
