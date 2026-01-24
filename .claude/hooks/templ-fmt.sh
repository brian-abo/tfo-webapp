#!/usr/bin/env bash
set -euo pipefail

# Skip if templ is not installed
command -v templ >/dev/null 2>&1 || exit 0

# Skip if no .templ files exist yet
if command -v git >/dev/null 2>&1 && git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  count="$(git ls-files '*.templ' | wc -l | tr -d ' ')"
else
  count="$(find . -name '*.templ' | wc -l | tr -d ' ')"
fi

[[ "$count" -gt 0 ]] || exit 0

templ fmt .

