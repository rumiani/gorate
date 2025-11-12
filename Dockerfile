# Dev Dockerfile
FROM golang:1.25-alpine

WORKDIR /app

# Install CompileDaemon
RUN apk add --no-cache git && \
    go install github.com/githubnemo/CompileDaemon@latest

# Copy go.mod/go.sum first for caching
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the files
COPY . .

# Run CompileDaemon to watch .go files
CMD CompileDaemon -build="go build -o gorate main.go" -command=./gorate -include="*.go"
