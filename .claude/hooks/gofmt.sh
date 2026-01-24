#!/usr/bin/env bash
set -euo pipefail
command -v go >/dev/null || exit 0
gofmt -w $(git ls-files '*.go' 2>/dev/null || true)

