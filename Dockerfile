##################################
FROM golang:1.21 as go-base
##################################

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
# Copy everything again (idk if this makes it slower or not, i hope someone can tell me if it does lol)
COPY . .
# -o specifies output location (bin folder, app executable)
RUN go build -o bin/app cmd/main.go

##################################
FROM node:20 AS node-base
##################################

WORKDIR /build
# apparently copying this first is faster?
COPY package*.json ./ 
RUN npm install
COPY . .

RUN npx tailwindcss init

COPY --from=go-base /build/bin/app /build/bin/app

EXPOSE 6969
#Where the containers main executable is
# ENTRYPOINT [ "/build/bin/app" ]


# WHAT TO DO:
# First add a .env file and specifiy LISTEN_ADDY

# COMMANDS TO RUN:
# docker build . -t lofi-docker:latest
# docker run lofi-docker:latest
