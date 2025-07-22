FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=8080 \
    GIN_MODE=release

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/bonsai ./cmd/bonsai/main.go

FROM alpine:latest

ENV GIN_MODE=release

WORKDIR /root/

COPY --from=builder /app/bonsai .

EXPOSE 8080

ENTRYPOINT ["./bonsai"]

