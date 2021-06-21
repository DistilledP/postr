package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/DistilledP/postr/internal/file"
	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/server"
	"github.com/DistilledP/postr/internal/util"
)

func usageInstructions() {
	fmt.Println(`Usage instructions:
===================

./image-upload <path_to_file>

To change the port:

GRPC_PORT=<port> ./image-upload <path_to_file>` + "\n")
}

func parseArgs(args []string) []*pb.ImageUpload {

	var result []*pb.ImageUpload
	for _, arg := range args {
		payload, err := os.ReadFile(arg)
		if err != nil {
			log.Printf("failed to load file %s: %v", arg, err)
			continue
		}

		result = append(result, &pb.ImageUpload{
			Name:     file.FileName(arg),
			Payload:  payload,
			FileType: file.MimeType(payload),
		})
	}

	return result
}

func uploadFile(ctx context.Context, client pb.PostrClient, upload *pb.ImageUpload) (*pb.ImageUploadResponse, error) {
	callOpts := []grpc.CallOption{}

	return client.Upload(ctx, upload, callOpts...)
}

const tpmlUploadResult = `
	Upload        : %v
	Filename      : %v
	FileType      : %v
	Size          : %d

`

const tpmlUploadResultWithError = `
	Error Message : %s

`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usageInstructions()
		os.Exit(1)
	}

	uploads := parseArgs(args)

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
		response, err := uploadFile(ctx, client, upload)

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

func displayResult(upload *pb.ImageUpload, result *pb.ImageUploadResponse) {
	switch result.Status {
	case pb.Status_FAILED:
		log.Printf(
			strings.TrimRight(tpmlUploadResult, "\n")+tpmlUploadResultWithError,
			result.Status,
			upload.Name,
			upload.FileType,
			result.SizeInBytes,
			result.ErrorMessage,
		)

	case pb.Status_SUCCESS:
		log.Printf(
			tpmlUploadResult,
			result.Status,
			upload.Name,
			upload.FileType,
			result.SizeInBytes,
		)
	}
}
