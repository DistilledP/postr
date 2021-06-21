package file

import (
	"os"
	"strings"

	"github.com/DistilledP/postr/internal/util"
)

func FileName(filePath string) string {
	fileParts := strings.Split(filePath, string(os.PathSeparator))
	fileName, _ := util.StringSlicePop(fileParts)

	return fileName
}
