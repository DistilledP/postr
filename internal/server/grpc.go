package server

import (
	"google.golang.org/grpc"

	"github.com/DistilledP/postr/internal/middleware"
	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/server/handler"
)

const DefaultGRPCPort = "3001"

func NewGRPCServer() *grpc.Server {
	var opts []grpc.ServerOption

	opts = append(opts, grpc.ChainUnaryInterceptor(middleware.UnaryServerMiddleware...))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPostrServer(grpcServer, &handler.GRPCHandler{})

	return grpcServer
}
