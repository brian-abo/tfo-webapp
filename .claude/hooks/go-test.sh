#!/usr/bin/env bash
set -euo pipefail
command -v go >/dev/null || exit 0
go test ./...
