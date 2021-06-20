package file

import (
	"net/http"

	pb "github.com/DistilledP/postr/internal/proto"
)

func MimeType(content []byte) pb.MimeType {
	switch http.DetectContentType(content) {
	case "image/jpeg":
		return pb.MimeType_JPEG

	case "image/gif":
		return pb.MimeType_GIF

	case "image/png":
		return pb.MimeType_PNG

	default:
		return pb.MimeType_UNKNOWN
	}
}

func IsAcceptedType(mimeType pb.MimeType) bool {
	return mimeType != pb.MimeType_UNKNOWN
}
