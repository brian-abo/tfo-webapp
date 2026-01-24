#!/usr/bin/env bash
set -euo pipefail

# Skip if not a Go module yet
[[ -f "go.mod" ]] || exit 0

command -v go >/dev/null || exit 0
gofmt -w $(git ls-files '*.go' 2>/dev/null || true)

