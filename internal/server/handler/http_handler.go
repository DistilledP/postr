package handler

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/DistilledP/postr/internal/file"
	"github.com/DistilledP/postr/internal/util"
)

type RootHandler struct {
}

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handle501(w)
		return
	}

	imageDir := util.GetImageDir()

	fileName := strings.TrimLeft(r.RequestURI, "/")
	filePath := fmt.Sprintf("%s/%s", imageDir, fileName)

	stat, err := os.Stat(filePath)
	if err != nil {
		handle404(w)
		return
	}

	content, err := file.ReadFile(imageDir, fileName)
	if err != nil {
		handle500(w, err)
		return
	}

	if !file.IsAcceptedType(file.MimeType(content)) {
		handle404(w)
		return
	}

	handle200(w, stat, content)
}

func handle404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("File doesn't exist"))
}

func handle500(w http.ResponseWriter, err error) {
	log.Printf("Server error: %v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

func handle501(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented"))
}

func handle200(w http.ResponseWriter, stat fs.FileInfo, content []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Length", string(rune(stat.Size())))
	w.Header().Add("Content-Type", http.DetectContentType(content))
	w.Write(content)
}
