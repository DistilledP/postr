package file

import (
	"fmt"
	"io/fs"
	"os"

	pb "github.com/DistilledP/postr/internal/proto"
)

func SaveFile(rootDir string, upload *pb.ImageUpload) (fs.FileInfo, error) {
	saveFilePath := fmt.Sprintf("%s/%s", rootDir, upload.Name)

	err := os.WriteFile(saveFilePath, upload.Payload, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write file: %v", err)
	}

	return os.Stat(saveFilePath)
}

func ReadFile(rootDir, fileName string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", rootDir, fileName))
}
