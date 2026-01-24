# tfo-webapp

A Go web application using the GOAT stack (Go, htmx/templ, Alpine.js, Tailwind CSS).

## Prerequisites

- Go 1.25+
- Node.js via [nvm](https://github.com/nvm-sh/nvm) (for Tailwind CSS)
- [templ](https://templ.guide/) for HTML templating
- [golangci-lint](https://golangci-lint.run/) for linting
- [goose](https://github.com/pressly/goose) for database migrations
- [tmux](https://github.com/tmux/tmux) for dev workflow
- Docker for local PostgreSQL

### Setup

```bash
# Node.js (via nvm)
nvm install --lts
nvm alias default 'lts/*'

# Go tools
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/pressly/goose/v3/cmd/goose@latest

# golangci-lint (macOS)
brew install golangci-lint
```

Ensure `$GOPATH/bin` is in your PATH.

## Development Workflow

```bash
# Start dev environment (tmux: server + css-watch)
make dev

# Or run components separately
make gen           # Run code generation (templ)
make css           # Build minified CSS
make css-watch     # CSS watch mode
go run ./cmd/tfo-webapp  # Run server
```

### Validation

```bash
make lint    # Run golangci-lint
make test    # Run tests (includes gen)
make verify  # Full validation: gen, css, lint, test
```

## Database

### Local PostgreSQL

```bash
make db-up    # Start Postgres container
make db-down  # Stop Postgres container
```

Connection: `postgres://postgres:postgres@localhost:5432/tfo_webapp_dev?sslmode=disable`

### Migrations

```bash
make migrate         # Run pending migrations
make migrate-down    # Roll back one migration
make migrate-status  # Show migration status

# Create new migration
goose -dir db/migrations create <name> sql
```

Migrations are stored in `db/migrations/`. See design docs for guidelines.
