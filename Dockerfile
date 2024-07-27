FROM golang:1.19.2-alpine3.16 AS builder

RUN apk add --no-cache git upx

# Prepare Git
RUN --mount=type=secret,id=GITHUB_USERNAME \
    --mount=type=secret,id=GITHUB_TOKEN \
    export GITHUB_USERNAME=$(cat /run/secrets/GITHUB_USERNAME) && \
    export GITHUB_TOKEN=$(cat /run/secrets/GITHUB_TOKEN) && \
    git config --global url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./

RUN go mod download

# Copy the entire project
COPY . .

# If you need to compile the package, you can do so with `go build` without the `-o` flag
RUN go build -v ./...
