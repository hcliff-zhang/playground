package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/hcliff-zhang/playground/server/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisterGRPCHandlers registers the API service implementation with the gRPC server.
// The service parameter should implement the serverpb.ApiServer interface.
func RegisterGRPCHandlers(grpcServer *grpc.Server, service pb.ApiServer) {
	pb.RegisterApiServer(grpcServer, service)
}

// RegisterHTTPGateway creates and registers a gRPC gateway handler that proxies
// HTTP requests to the gRPC server running on the specified port.
// It returns an HTTP handler that can be used to serve the gateway.
func RegisterHTTPGateway(ctx context.Context, grpcPort string) (http.Handler, error) {
	// Create a new gRPC gateway multiplexer
	mux := runtime.NewServeMux()

	// Set up a connection to the gRPC server
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcEndpoint := fmt.Sprintf("localhost%s", grpcPort)

	// Register the service handler
	err := pb.RegisterApiHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway handler: %w", err)
	}

	return mux, nil
}
