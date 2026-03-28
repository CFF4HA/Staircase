# This is the first stage environment.
FROM golang:latest

WORKDIR /build
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o web cmd/web/main.go


