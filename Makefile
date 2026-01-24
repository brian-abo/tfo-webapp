.PHONY: css css-watch migrate migrate-down migrate-status

# Database connection string (override with environment variable)
DATABASE_URL ?= postgres://localhost:5432/tfo_webapp?sslmode=disable

css:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --minify

css-watch:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --watch

migrate:
	goose -dir db/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir db/migrations postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir db/migrations postgres "$(DATABASE_URL)" status
