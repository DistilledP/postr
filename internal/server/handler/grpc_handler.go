package handler

import (
	"context"
	"log"

	pb "github.com/DistilledP/postr/internal/proto"
)

type GRPCHandler struct {
	pb.UnimplementedPostrServer
}

func (s GRPCHandler) Upload(_ context.Context, upload *pb.ImageUpload) (*pb.ImageUploadResponse, error) {
	log.Println(upload.Name)

	return &pb.ImageUploadResponse{
		Status:      pb.Status_SUCCESS,
		SizeInBytes: 1000,
	}, nil
}
