package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/server"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	grpcAddr := server.GetAddress("GRPC_PORT", server.DefaultGRPCPort)
	conn, err := grpc.DialContext(ctx, grpcAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial %s: %v", grpcAddr, err)
	}
	defer conn.Close()

	client := pb.NewPostrClient(conn)

	upload := &pb.ImageUpload{
		Name:     "test.cat",
		FileType: pb.MimeType_GIF,
		Payload:  []byte("")}

	callOpts := []grpc.CallOption{}

	response, err := client.Upload(ctx, upload, callOpts...)
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}

	log.Println(response.Status, response.ErrorMessage, response.SizeInBytes)
}
