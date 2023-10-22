# A multi-build stage for Go-Application

# Build Stage
FROM golang:1.15 AS builder

WORKDIR /app

COPY . .

RUN go build -o hello-golang

# Final Stage
FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/hello-golang .

RUN chmod +x hello-golang

CMD ["./hello-golang"]
