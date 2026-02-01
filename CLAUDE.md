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
- Functions that perform I/O or may block must accept `context.Context` as the first parameter
- Never instantiate a context inside a library function; always accept it from the caller
- All tests use Gomega (dot-import `github.com/onsi/gomega`, `g := NewWithT(t)`)
- Tests must use `t.Context()` instead of `context.Background()`
- Validate all external input at system boundaries (user input, HTTP requests, external APIs)
- Use `database/sql` parameterized queries (`$1`, `$2`) for all SQL -- never string concatenation
- Use `net/mail.ParseAddress` for email validation on server-side form handling

## Generated files policy
- templ-generated Go files (`*_templ.go`) are NOT committed.
- `.templ` files are the source of truth.
- The build/test pipeline must run `templ generate` before `go test`.
- Do not attempt to stage or commit generated templ output.

## Git + atomic commits
- Commits must be small and atomic (one logical change per commit).
- Each commit must compile and tests should pass, or explicitly state why they cannot.
- Prefer Conventional Commits:
  `type(scope): subject`
- Reference Beads issues when applicable:
  `Refs: BD-<id>`
- Use `/commit` to propose or create commits.

## PR-only workflow
- All code changes must go through a Pull Request (no direct commits to main).
- Create a feature branch for each bead or logical change set.
- Use `/create-pr` to draft/open PRs.

## Required skills
These skills MUST be used for their respective operations. Do not fall back to default behaviors.

- `/commit` - MUST be used for all git commits (overrides default commit behavior)
- `/create-pr` - MUST be used for all pull requests
- `/go-change` - Use for non-trivial Go implementation work
- `/verify-ui` - MUST be run before committing changes that affect observable UI

## UI verification workflow
Before committing any changes that could affect what the user sees in the browser:
1. Run `/verify-ui` to build assets and start the server
2. Wait for user to visually inspect and approve
3. Only proceed to `/commit` after approval

This includes (but is not limited to):
- Templates and components (`web/`)
- Handlers that render pages
- Models/data that surface in the UI
- Database migrations affecting displayed data

## Work tracking (Beads)
- Beads (`bd`) is the **only** task / issue system.
- Do NOT create TODO.md / PLAN.md / TASKS.md.
- Do NOT expose email address in tasks.
- Before starting work: review Beads for active issues.
- After finishing work: update the issue with results and validation evidence.

## Bead phase labels
Beads use `phase:<value>` labels to track readiness:

- `phase:planning` -- has open questions or unresolved design decisions. NOT ready for implementation.
- `phase:ready` -- planned, questions resolved, ready to pick up.
- No phase label -- legacy beads or already completed.

Rules:
- Do NOT start implementation on a `phase:planning` bead. Resolve questions first.
- When all questions on a bead are answered and acceptance criteria are clear, swap label: `bd label remove <id> phase:planning && bd label add <id> phase:ready`
- Use `bd list --label phase:planning` to see what needs planning.
- Use `bd list --label phase:ready` to see what's ready for work.
- `bd ready` shows unblocked beads regardless of phase. Check phase labels before starting.

## Bead structure contract
Unless explicitly stated otherwise, Beads should follow this structure:

- Problem / Goal
- Context / Constraints (only if relevant)
- Decisions
- Risks
- Acceptance criteria
- If information is missing or ambiguous, ask clarifying questions before updating the bead.
- Once updated, the bead should be ready for implementation.
- Avoid "Open Questions" sections inside beads.

Not all sections are required for every bead, but Problem and Acceptance
should always be present. Keep sections concise.

## Evidence required in responses
When code changes:
- Files changed
- Commands run
- Outcomes (pass/fail)
- Known gaps, risks, or follow-ups
