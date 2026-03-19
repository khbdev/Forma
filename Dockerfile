FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o forma ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/forma .

EXPOSE 8081

CMD ["./forma"]