package main

import (
	"fmt"
	"log"
	"strings"

	pb "github.com/DistilledP/postr/internal/proto"
)

func usageInstructions() {
	fmt.Println(`Usage instructions:
===================

./image-upload <path_to_file>

To change the port:

GRPC_PORT=<port> ./image-upload <path_to_file>` + "\n")
}

const templateUploadResult = `
	Upload        : %v
	Filename      : %v
	FileType      : %v
	Size          : %d

`

const templateUploadResultWithError = `
	Error Message : %s

`

func displayResult(upload *pb.ImageUpload, result *pb.ImageUploadResponse) {
	switch result.Status {
	case pb.Status_FAILED:
		log.Printf(
			strings.TrimRight(templateUploadResult, "\n")+templateUploadResultWithError,
			result.Status,
			upload.Name,
			upload.FileType,
			result.SizeInBytes,
			result.ErrorMessage,
		)

	case pb.Status_SUCCESS:
		log.Printf(
			templateUploadResult,
			result.Status,
			upload.Name,
			upload.FileType,
			result.SizeInBytes,
		)
	}
}
