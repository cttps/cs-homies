run: build
	@./bin/app

build:
	@go build -o bin/app cmd/main.go

css:
	npx tailwindcss -i frontend/static/input.css -o frontend/static/output.css --watch

FROM golang:1.21 as goPart

WORKDIR /build
# apparently copying this first is faster?
COPY package*.json ./ 

COPY . .
RUN go mod download

# -o specifies output location (we make a bin folder in build)
RUN go build -o bin/app cmd/main.go

EXPOSE 6969

#Where the containers main executable is
ENTRYPOINT [ "/build/bin/app" ]


# WHAT TO DO:
# First add a .env file and specifiy LISTEN_ADDY

# COMMANDS TO RUN:
# docker build . -t lofi-docker:latest
# docker run lofi-docker:latest
# docker run tailwindcss -i frontend/static/input.css -o frontend/static/output.css --watch lofi-docker:latest
