run: build
	@./bin/app

build:
	@go build -o bin/app cmd/main.go

css:
	npx tailwindcss -i frontend/static/input.css -o frontend/static/output.css --watch