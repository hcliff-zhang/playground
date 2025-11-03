# Multi-stage Dockerfile for the Go "playground" application
# Builder stage
FROM golang:1.20-alpine AS builder

WORKDIR /src

# Install git for module downloads (some modules require git)
RUN apk add --no-cache git

# Cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source and build a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags='-s -w' -o /out/playground ./main.go

# Final image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

# Create non-root user
RUN addgroup -S app && adduser -S -G app app

COPY --from=builder /out/playground /usr/local/bin/playground
RUN chown app:app /usr/local/bin/playground

USER app

EXPOSE 8080 9090

ENTRYPOINT ["/usr/local/bin/playground"]
