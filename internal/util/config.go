package util

import (
	"log"
	"os"
	"strings"
)

const DefaultImageDir = "/tmp"
const EnvImageDirKey = "IMAGE_DIR"

func GetImageDir() string {
	imageDir := strings.TrimSpace(os.Getenv(EnvImageDirKey))
	if imageDir == "" {
		imageDir = DefaultImageDir
	}

	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.Mkdir(imageDir, 0644)
		if err != nil {
			log.Printf("Failed to make dir %s, %v - using %s\n", imageDir, err, DefaultImageDir)
		}
		imageDir = DefaultImageDir
	}

	return imageDir
}

func GetAddress(key, defaultValue string) string {
	address := defaultValue
	if envAddress := strings.TrimSpace(os.Getenv(key)); envAddress != "" {
		address = envAddress
	}

	if !strings.Contains(address, ":") {
		address = ":" + address
	}

	return address
}
