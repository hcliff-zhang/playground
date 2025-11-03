package application

import (
	pb "github.com/hcliff-zhang/playground/server/serverpb"
	"google.golang.org/grpc"
)

// RegisterGRPCHandlers registers the API service implementation with the gRPC server.
// The service parameter should implement the serverpb.ApiServer interface.
func RegisterGRPCHandlers(grpcServer *grpc.Server, service pb.ApiServer) {
	pb.RegisterApiServer(grpcServer, service)
}
