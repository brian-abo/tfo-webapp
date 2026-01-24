.PHONY: css css-watch

css:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --minify

css-watch:
	npx @tailwindcss/cli -i web/static/input.css -o web/static/styles.css --watch
