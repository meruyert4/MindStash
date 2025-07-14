# --- Stage 1: Build ---
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o mindstash ./cmd/server

# --- Stage 2: Run ---
FROM alpine:latest

RUN adduser -D -g '' mindstashuser

WORKDIR /app
COPY --from=builder /app/mindstash .
COPY --from=builder /app/frontend ./frontend

USER mindstashuser

EXPOSE 8080

CMD ["./mindstash"]
