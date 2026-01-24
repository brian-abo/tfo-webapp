.PHONY: dev css css-watch templ gen lint test verify migrate migrate-down migrate-status db-up db-down

# Database connection string (override with environment variable)
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/tfo_webapp_dev?sslmode=disable

# Development
dev: gen
	tmux new-session -d -s tfo-webapp 'go run ./cmd/tfo-webapp' \; \
		split-window -h 'make css-watch' \; \
		attach

# Code generation
templ:
	templ generate

gen: templ

# CSS
css:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --minify

css-watch:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --watch

# Validation
lint:
	golangci-lint run ./...

test: gen
	go test ./...

verify: gen css lint test

# Database migrations
migrate:
	goose -dir db/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir db/migrations postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir db/migrations postgres "$(DATABASE_URL)" status

# Local database
db-up:
	docker compose up -d db

db-down:
	docker compose down
