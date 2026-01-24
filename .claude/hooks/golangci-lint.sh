#!/usr/bin/env bash
set -euo pipefail
command -v golangci-lint >/dev/null || exit 0
golangci-lint run ./...
