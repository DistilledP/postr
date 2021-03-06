package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/DistilledP/postr/internal/proto"
)

func TestMimeType(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected pb.MimeType
	}{
		{"Empty", []byte{}, pb.MimeType_UNKNOWN},
		{"Plain text", []byte(`This is not HTML. It has ☃ though.`), pb.MimeType_UNKNOWN},
		{"BMP Image", []byte("BM..."), pb.MimeType_UNKNOWN},
		{"GIF 87a", []byte(`GIF87a`), pb.MimeType_GIF},
		{"GIF 89a", []byte(`GIF89a...`), pb.MimeType_GIF},
		{"WEBP image", []byte("RIFF\x00\x00\x00\x00WEBPVP"), pb.MimeType_UNKNOWN},
		{"PNG image", []byte("\x89PNG\x0D\x0A\x1A\x0A"), pb.MimeType_PNG},
		{"JPEG image", []byte("\xFF\xD8\xFF"), pb.MimeType_JPEG},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MimeType(tc.data)

			assert.Equal(
				t,
				tc.expected,
				actual,
				fmt.Sprintf("test case %v failed: got %v, want %v", tc.name, actual, tc.expected),
			)
		})
	}
}

func TestIsAcceptedMimeType(t *testing.T) {
	testCases := []struct {
		name     string
		testType pb.MimeType
		expected bool
	}{
		{"Unknown", pb.MimeType_UNKNOWN, false},
		{"GIF", pb.MimeType_GIF, true},
		{"JPEG", pb.MimeType_JPEG, true},
		{"PNG", pb.MimeType_PNG, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsAcceptedType(tc.testType)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
