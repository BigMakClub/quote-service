FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o quote-service ./cmd/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/quote-service .
COPY ./db ./db

EXPOSE 8080
CMD ["./quote-service"]