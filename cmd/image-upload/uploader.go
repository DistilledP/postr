package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/DistilledP/postr/internal/proto"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock())

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		log.Fatal("Failed to dial :3000")
	}
	defer conn.Close()

	client := pb.NewPostrClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	upload := &pb.ImageUpload{
		Name:     "test.cat",
		FileType: pb.MimeType_GIF,
		Payload:  []byte("")}

	response, err := client.Upload(ctx, upload)
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}

	log.Println(response.Status, response.ErrorMessage, response.SizeInBytes)
}
