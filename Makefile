run: build
	@./bin/app

build:
	@go build -o bin/app cmd/main.go

docker-run:
	@docker-compose up --build

css:
	@npx tailwindcss -i frontend/static/input.css -o frontend/static/output.css --watch

css-compile:
	@npx tailwindcss -i frontend/static/input.css -o frontend/static/output.css --minify

test:
	@echo You have make so dw