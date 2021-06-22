package util

import (
	"log"
	"os"

	"github.com/DistilledP/postr/internal/file"
	pb "github.com/DistilledP/postr/internal/proto"
)

func ParseArgs(args []string) []*pb.ImageUpload {
	var result []*pb.ImageUpload
	for _, arg := range args {
		payload, err := os.ReadFile(arg)
		if err != nil {
			log.Printf("failed to load file %s: %v", arg, err)
			continue
		}

		result = append(result, &pb.ImageUpload{
			Name:     FileName(arg),
			Payload:  payload,
			FileType: file.MimeType(payload),
		})
	}

	return result
}
