# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY ./api ./api
RUN mkdir -p ./pkg/oas
RUN go generate ./...

COPY . ./

RUN CGO_ENABLED=0 go build -o /app-server ./cmd/app


FROM alpine:latest

ENV SERVICE_NAME="shedevr-backend"

RUN addgroup --gid 1001 -S $SERVICE_NAME && \
    adduser -G $SERVICE_NAME --shell /bin/false --disabled-password -H --uid 1001 $SERVICE_NAME && \
    mkdir -p /var/log/$SERVICE_NAME && \
    chown $SERVICE_NAME:$SERVICE_NAME /var/log/$SERVICE_NAME


COPY --from=builder /app-server /app/

WORKDIR /app

RUN chmod u+x app-server

COPY ./config/config.example.env ./config/config.env

ENV SHEDEVR_BACKEND_API_PORT=8080
EXPOSE 8080

ENV SHEDEVR_BACKEND_METRICS_PORT=8090
EXPOSE 8090

RUN apk update && apk add curl

USER $SERVICE_NAME

HEALTHCHECK --interval=5s --timeout=10s --retries=3 CMD curl -sS 127.0.0.1:8080 || exit 1

CMD ["./app-server"]