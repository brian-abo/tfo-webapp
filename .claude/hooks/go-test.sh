
#!/usr/bin/env bash
set -euo pipefail

# Skip if not a Go module yet
[[ -f "go.mod" ]] || exit 0

go test ./...
