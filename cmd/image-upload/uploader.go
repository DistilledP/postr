package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/server"
	"github.com/DistilledP/postr/internal/util"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usageInstructions()
		os.Exit(1)
	}

	uploads := util.ParseArgs(args)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	grpcAddr := util.GetAddress("GRPC_PORT", server.DefaultGRPCPort)
	conn, err := grpc.DialContext(ctx, grpcAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial %s: %v", grpcAddr, err)
	}
	defer conn.Close()

	client := pb.NewPostrClient(conn)

	var successCount int
	for _, upload := range uploads {
		response, err := util.UploadFile(ctx, client, upload)

		if err != nil {
			log.Printf("failed to upload file: %s, error: %v", upload.Name, err)
			continue
		}

		if response.Status == pb.Status_SUCCESS {
			successCount++
		}

		displayResult(upload, response)
	}

	log.Printf("Uploaded %d out of %d files\n", successCount, len(uploads))
}
