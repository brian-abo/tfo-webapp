# Claude Code – Project Instructions (GOAT Stack)

## Engineering posture
- Maintainability > cleverness
- Small, reviewable diffs
- No guessing: ask targeted questions if unclear
- Tests are mandatory and a part of the definition of done

## Mandatory workflow
1. Explore the existing codebase and established patterns
2. Propose a short plan (files + approach)
3. Implement incrementally in small diffs
4. Verify via tools (format, lint, tests)
5. Summarize changes and validation results

## Design principles
- Each module does **one** thing well (Single Responsibility)
- Code should be modular with strict separation of concerns
- Related functions belong together
- Minimize dependencies between modules using interfaces/abstractions
- Do not introduce abstraction without clear justification

Guiding principles (apply pragmatically, not dogmatically):
DRY, KISS, YAGNI, SOLID, Law of Demeter, Open/Closed Principle

## Communication style
- Be concise. Write as if communicating with a senior engineer.
- Prefer bullets over prose.
- Explanations: ~2–5 sentences unless more depth is explicitly requested.
- No filler, no narration, no motivational language.
- **No emojis in conversation output.**

## Codebase hygiene
- Do NOT add meta files (TODO.md, NOTES.md, INSTRUCTIONS.md, PLAN.md, etc.).
- Tracking work happens in Beads only (`bd`).
- Comments are rare and intentional:
  - Allowed only for non-obvious invariants, tricky logic, or sharp edges
  - Prefer clear naming and structure over comments
- **No emojis in file names, code, commit messages, or documentation.**
  - Exception: emojis are allowed in scripts or CLI output *only* when they materially improve UX.

## Stack (GOAT)
We use the GOAT stack:
- Go backend
- Server-rendered UI via `templ`
- Alpine.js for light, localized interactivity
- Tailwind CSS for styling

Constraints:
- No heavy frontend frameworks (React, Vue, etc.)
- Minimal JavaScript; prefer HTML + templ + Alpine
- Follow official tooling and conventions (templ fmt/generate, Tailwind build)

## Go standards
- gofmt is mandatory
- golangci-lint must pass
- Tests required for behavior changes
- Errors wrapped intentionally using `%w`
- Public APIs require godoc comments

## Git + atomic commits
- Commits must be small and atomic (one logical change per commit).
- Each commit must compile and tests should pass, or explicitly state why they cannot.
- Prefer Conventional Commits:
  `type(scope): subject`
- Reference Beads issues when applicable:
  `Refs: BD-<id>`
- Use `/commit` to propose or create commits.
- Use `/create-pr` to draft or open pull requests.

## Work tracking (Beads)
- Beads (`bd`) is the **only** task / issue system.
- Do NOT create TODO.md / PLAN.md / TASKS.md.
- Before starting work: review Beads for active issues.
- After finishing work: update the issue with results and validation evidence.

## Evidence required in responses
When code changes:
- Files changed
- Commands run
- Outcomes (pass/fail)
- Known gaps, risks, or follow-ups
