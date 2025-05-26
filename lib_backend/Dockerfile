FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLE=0 go build -o /app/bin/server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/server ./server

EXPOSE 8080

ENV GIN_MODE=release

ENTRYPOINT ["./server"]