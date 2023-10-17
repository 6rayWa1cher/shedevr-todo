# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/db-migrate ./cmd/db-migrate

RUN CGO_ENABLED=0 go build -o ./db-migrate ./cmd/db-migrate

RUN chmod +x ./db-migrate

CMD ["./db-migrate"]
