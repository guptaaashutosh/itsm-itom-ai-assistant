# Dockerfile for ITSM-ITOM AI Assistant
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o itsm-app ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/itsm-app .
COPY .config/.env ./.env
EXPOSE 8080
CMD ["./itsm-app"] 