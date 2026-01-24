---
name: create-pr
description: Draft or open a PR with a structured description, evidence, and Beads references.
---

## When to use
- User asks to open a PR / create PR / “make a PR for this”
- Or invoke with `/create-pr`

## Branch naming
- Pattern: `bd-<bead-id>-<short-description>`
- Example: `bd-tfo-webapp-98w-user-model`

## Workflow
0) **Branch validation (FIRST)**:
   - Run `git branch --show-current`
   - If on `main` or `master`: STOP
     - "Cannot create PR from main. Use `/commit` first to set up a feature branch."
   - Validate branch name matches `bd-<bead-id>-*` pattern (warn if not)
   - Verify commits exist ahead of main: `git rev-list --count main..HEAD`
     - If 0 commits: STOP, nothing to PR

1) Gather context:
   - Current branch + target branch
   - Commits included (`git log --oneline <target>..HEAD`)
   - High-level diff (`git diff <target>...HEAD --stat`)
2) Produce PR content (Markdown):
   - Title (Conventional-ish, human readable)
   - Summary (3–6 bullets)
   - Changes (bullets grouped by area)
   - Testing (commands + results)
   - UI notes (if applicable)
   - Beads references: `Refs: BD-<id>` list
   - Risks/rollout notes (only if meaningful)
3) If user wants to actually open the PR:
   - Prefer `gh pr create` if available
   - Otherwise output the exact title/body ready to paste

## Output style
- Be concise. Engineer-readable. No fluff.

## Notes
- If the PR mixes concerns, stop and propose splitting.

