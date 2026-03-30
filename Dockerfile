# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o api cmd/api/main.go
RUN go build -o worker cmd/worker/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api .
COPY --from=builder /app/worker .

EXPOSE 8080

CMD ["./api"]