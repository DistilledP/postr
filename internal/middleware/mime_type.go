package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/DistilledP/postr/internal/file"
	pb "github.com/DistilledP/postr/internal/proto"
)

func mimeTypeCheckInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if imageUpload, ok := req.(*pb.ImageUpload); ok {
		if resp := imageUploadMimeResponse(imageUpload); resp != nil {
			return resp, nil
		}
	}

	return handler(ctx, req)
}

func imageUploadMimeResponse(imageUpload *pb.ImageUpload) *pb.ImageUploadResponse {
	if !file.IsAcceptedType(file.MimeType(imageUpload.GetPayload())) {
		log.Println("failed mime check")

		return &pb.ImageUploadResponse{
			Status:       pb.Status_FAILED,
			SizeInBytes:  0,
			ErrorMessage: fmt.Sprintf(`File format is not accepted, detected: "%s"`, http.DetectContentType(imageUpload.Payload)),
		}
	}

	return nil
}
