FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/api

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8000
CMD ["./server"]