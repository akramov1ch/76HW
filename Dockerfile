FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o hello .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/hello .
COPY .env .

EXPOSE 9000
CMD ["./hello"]
