FROM golang:1.21

WORKDIR /build
COPY . .
RUN go mod download

# -o specifies output location (we make a bin folder in build)
RUN go build -o bin/app cmd/main.go

#Where the containers main executable is
ENTRYPOINT [ "/build/bin/app" ]


# COMMANDS TO RUN:
# docker build . -t lofi-docker:latest
# docker run -e LISTEN_PORT=6969 -p 6969:6969 lofi-docker:latest