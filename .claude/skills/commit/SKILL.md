---
name: commit
description: Create atomic, conventional commits with Beads references and validation evidence.
---

## Principles
- Atomic: one logical change per commit.
- Each commit should build and tests should pass (or explicitly explain why not).
- Prefer Conventional Commits: type(scope): subject
- Reference Beads issues when relevant: `Refs: BD-<id>` (footer)

## Workflow
1) Inspect repo state:
   - `git status`
   - `git diff`
   - `git diff --staged`
2) Propose a commit plan:
   - Identify logical change groups (1–3 max)
   - For each group, propose:
     - files involved
     - commit message (type(scope): subject)
     - whether it needs tests updated/added
3) Stage *surgically*:
   - Prefer `git add -p` (or file-by-file) over `git add .`
4) Validate before committing (unless user says otherwise):
   - gofmt / templ fmt already handled by hooks
   - `golangci-lint run ./...` (if available)
   - `go test ./...`
5) Commit:
   - Use the proposed message
   - Include footer when applicable:
     - `Refs: BD-<id>`
6) Output (be concise):
   - Commit hash(es)
   - Messages
   - Commands run + outcomes
   - Any follow-ups

## Safety / constraints
- Never commit secrets.
- Never change unrelated files just to “clean up.”
- If changes are too mixed, stop and propose a split before committing.

## Examples
- `feat(auth): add login handler`
- `fix(db): handle nil tx`
- `refactor(ui): extract templ component`

