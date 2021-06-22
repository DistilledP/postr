package util

import (
	"os"
	"strings"
)

func FileName(filePath string) string {
	fileParts := strings.Split(filePath, string(os.PathSeparator))
	fileName, _ := StringSlicePop(fileParts)

	return fileName
}
