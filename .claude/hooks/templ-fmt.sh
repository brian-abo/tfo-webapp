#!/usr/bin/env bash
set -euo pipefail

# Only run if templ exists AND there are any .templ files
command -v templ >/dev/null 2>&1 || exit 0

if command -v git >/dev/null 2>&1 && git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  count="$(git ls-files '*.templ' | wc -l | tr -d ' ')"
else
  count="$(find . -name '*.templ' | wc -l | tr -d ' ')"
fi

if [[ "${count}" == "0" ]]; then
  exit 0
fi

templ fmt .

