package util

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/DistilledP/postr/internal/proto"
)

func UploadFile(ctx context.Context, client pb.PostrClient, upload *pb.ImageUpload) (*pb.ImageUploadResponse, error) {
	callOpts := []grpc.CallOption{}

	return client.Upload(ctx, upload, callOpts...)
}
