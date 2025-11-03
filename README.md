# Hello World (Go)

This is a minimal Go "Hello, world!" example.

Files:

- `main.go` - contains the `Hello()` function and `main` which prints the greeting.
- `hello_test.go` - simple unit test for `Hello()`.

How to run locally:

1. Ensure you have Go installed (go version >= 1.20 recommended).
2. From the project root (`/Users/heathcliff/playground`):

```bash
# run tests
go test ./...

# build (creates an executable named `playground`)
go build -o playground

# run the program
./playground
```

That's it — a tiny verified example.


Usage (attach to server):

```go
import "google.golang.org/grpc"

srv := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingUnaryServerInterceptor()),
)
```

Run tests to verify the interceptors:

```bash
go test ./...
```

Docker
------

A multi-stage `Dockerfile` is included to build the Go binary and produce a small runtime image.

Build the image from the project root:

```bash
docker build -t playground:latest .
```

Run the container (exposes HTTP 8080 and gRPC 9090 by default):

```bash
docker run --rm -p 8080:8080 -p 9090:9090 playground:latest
```

Notes:
- The Dockerfile uses an Alpine-based builder and runtime. It builds for linux/amd64.
- If you need a different architecture or want to include TLS files or config, mount them with `-v` or extend the Dockerfile.
