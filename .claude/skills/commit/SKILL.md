---
name: commit
description: Create atomic, conventional commits with Beads references and validation evidence.
---

## Principles
- Atomic: one logical change per commit.
- Commit message MUST be: type(scope): subject
  - scope is required (no exceptions).
- Each commit should build and tests should pass (or explicitly explain why not).
- Prefer Conventional Commits: type(scope): subject
- Reference Beads issues when relevant: `Refs: BD-<id>` (footer)

## Commit scope rules
- Commit messages MUST follow: type(scope): subject
- Scope is REQUIRED.
- Allowed scopes:
  scaffold, web, domain, repo, auth, config,
  infra, ci, docs, test, refactor, chore
- If a change spans multiple scopes, stop and propose splitting the commit.
- Do not invent new scopes.

## Branch naming
- Pattern: `bd-<bead-id>-<short-description>`
- Example: `bd-tfo-webapp-98w-user-model`

## Workflow
0) **Branch setup (FIRST)**:
   - Run `git branch --show-current`
   - If on `main` or `master`:
     a. Pull latest: `git pull origin main`
     b. Ask user for bead ID and short description
     c. Create branch: `git checkout -b bd-<bead-id>-<description>`
     d. Confirm with user before proceeding
   - If on feature branch:
     - Verify branch name matches pattern (warn if not)
     - Check if behind main: `git fetch origin && git rev-list --count HEAD..origin/main`
     - If behind, warn user and suggest rebasing

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
4) Validate (unless user says otherwise):
   - formatting handled by hooks
   - `golangci-lint run ./...` (if available and go.mod exists)
   - `go test ./...` (if go.mod exists)
5) Commit:
   - Use the proposed message exactly.
   - Add Beads reference footer when applicable:
     - `Refs: <bead-id>`
   - DO NOT add any AI attribution lines:
     - no "Co-authored-by"
     - no "Generated-by"
     - no "AI-assisted"
   - Ensure git author is the user's normal git identity (do not override author).
   - Validate the commit message matches type(scope): subject with an allowed scope before committing.
6) Output (be concise):
   - Commit hash
   - Messages
   - Commands run + outcomes

## Safety / constraints
- Never commit secrets.
- Never change unrelated files just to “clean up.”
- If changes are too mixed, stop and propose a split before committing.

## Examples
- `feat(auth): add login handler`
- `fix(db): handle nil tx`
- `refactor(ui): extract templ component`

