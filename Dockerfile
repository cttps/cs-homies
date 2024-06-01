##################################
FROM golang:1.21 as go-base
##################################

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
# Copy everything (idk if everything it copies is necessary, i hope someone can tell me if it isnt lol)
COPY . .
RUN go build -o bin/app cmd/main.go

#Where the containers main executable is
ENTRYPOINT [ "/build/bin/app" ]



### COMMANDS TO RUN: ###
# make docker-run # OR # docker-compose up --build
# THEN
# make css (to watch tailwind realtime)