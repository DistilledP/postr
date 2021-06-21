package handler

import (
	"context"

	"github.com/DistilledP/postr/internal/file"
	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/util"
)

type GRPCHandler struct {
	pb.UnimplementedPostrServer
}

func (s GRPCHandler) Upload(_ context.Context, upload *pb.ImageUpload) (*pb.ImageUploadResponse, error) {
	stats, err := file.SaveFile(util.GetImageDir(), upload)

	if err != nil {
		return &pb.ImageUploadResponse{
			Status:       pb.Status_FAILED,
			SizeInBytes:  0,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ImageUploadResponse{
		Status:      pb.Status_SUCCESS,
		SizeInBytes: stats.Size(),
	}, nil
}
