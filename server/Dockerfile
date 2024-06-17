# Build stage
FROM golang:1.21.5 as builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
