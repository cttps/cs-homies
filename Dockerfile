FROM golang:1.21

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


# COMMANDS TO RUN:
# docker build . -t lofi-docker:latest
# docker run lofi-docker:latest