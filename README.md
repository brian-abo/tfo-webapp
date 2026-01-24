# tfo-webapp

A Go web application using the GOAT stack (Go, htmx/templ, Alpine.js, Tailwind CSS).

## Prerequisites

- Go 1.25+
- Node.js (for Tailwind CSS)
- [goose](https://github.com/pressly/goose) for database migrations
- PostgreSQL 16+ (local development via Docker)

### Installing goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Ensure `$GOPATH/bin` is in your PATH.

## Database Migrations

Migrations are managed with goose and stored in `db/migrations/`.

### Configuration

Set the database connection string via environment variable:

```bash
export DATABASE_URL="postgres://localhost:5432/tfo_webapp?sslmode=disable"
```

Or pass it directly to make commands (see below).

### Migration Commands

```bash
# Run all pending migrations
make migrate

# Roll back the most recent migration
make migrate-down

# Show migration status
make migrate-status
```

### Creating New Migrations

```bash
goose -dir db/migrations create <migration_name> sql
```

This creates a new SQL migration file with up and down sections:

```sql
-- +goose Up
-- SQL statements for applying the migration

-- +goose Down
-- SQL statements for reverting the migration
```

### Migration Guidelines

- One migration per logical change
- Always include both up and down migrations
- Use standard SQL compatible with PostgreSQL 16 and Aurora PostgreSQL
- Test migrations locally before deploying
- Avoid destructive operations (DROP, TRUNCATE) without explicit approval

## Development

### Build CSS

```bash
make css        # Build minified CSS
make css-watch  # Watch mode for development
```

### Run the Application

```bash
go run ./cmd/tfo-webapp
```
