package middleware

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/DistilledP/postr/internal/file"
	pb "github.com/DistilledP/postr/internal/proto"
)

func mimeMismatchInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if imageUpload, ok := req.(*pb.ImageUpload); ok {
		if resp := imageUploadMimeCheckResponse(imageUpload); resp != nil {
			return resp, nil
		}
	}

	return handler(ctx, req)
}

func imageUploadMimeCheckResponse(imageUpload *pb.ImageUpload) *pb.ImageUploadResponse {
	uploadMimeType := file.MimeType(imageUpload.Payload)

	if imageUpload.FileType != uploadMimeType {
		log.Println("failed mime mismatch")

		return &pb.ImageUploadResponse{
			Status:       pb.Status_FAILED,
			SizeInBytes:  0,
			ErrorMessage: fmt.Sprintf(`File format mis-match, expecting "%v", got "%v"`, imageUpload.FileType, uploadMimeType),
		}
	}

	return nil
}
