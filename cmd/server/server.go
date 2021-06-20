package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/DistilledP/postr/internal/middleware"
	pb "github.com/DistilledP/postr/internal/proto"
)

type Server struct {
	pb.UnimplementedPostrServer
}

func (s Server) Upload(_ context.Context, upload *pb.ImageUpload) (*pb.ImageUploadResponse, error) {
	log.Println(upload.Name)

	return &pb.ImageUploadResponse{
		Status:      pb.Status_SUCCESS,
		SizeInBytes: 1000,
	}, nil
}

func newServer() *Server {
	return &Server{}
}

func main() {
	sock, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen on port: 3000")
	}

	var opts []grpc.ServerOption

	opts = append(opts, grpc.ChainUnaryInterceptor(middleware.UnaryServerMiddleware...))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPostrServer(grpcServer, newServer())
	grpcServer.Serve(sock)
}
