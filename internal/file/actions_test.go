package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	pb "github.com/DistilledP/postr/internal/proto"
	"github.com/DistilledP/postr/internal/test"
)

func TestSaveFile(t *testing.T) {
	testCases := []struct {
		name        string
		rootDir     string
		upload      *pb.ImageUpload
		expected    fs.FileInfo
		expectedErr error
	}{
		{
			name:    "Writes file successfully",
			rootDir: "/tmp",
			upload: &pb.ImageUpload{
				Name:    "test1.txt",
				Payload: []byte("Test1"),
			},
			expected: test.NewMockFileInfo(
				"test1.txt",
				5,
				0644,
				time.Time{},
				false,
				nil,
			),
			expectedErr: nil,
		},
		{
			name:    "Failed to write file",
			rootDir: "/actions_test",
			upload: &pb.ImageUpload{
				Name:    "test2.txt",
				Payload: []byte("Test2"),
			},
			expected:    nil,
			expectedErr: errors.New("failed to write file: open /actions_test/test2.txt: no such file or directory"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SaveFile(tc.rootDir, tc.upload)

			assert.Equal(t, tc.expectedErr, err)

			if tc.expected != nil {
				assert.Equal(t, tc.expected.Name(), actual.Name())
				assert.Equal(t, tc.expected.Size(), actual.Size())
				assert.Equal(t, tc.expected.Mode(), actual.Mode())
			}
		})
	}

	t.Cleanup(func() {
		for _, tc := range testCases {
			os.Remove(fmt.Sprintf("%s/%s", tc.rootDir, tc.upload.Name))
		}
	})
}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name        string
		rootDir     string
		fileName    string
		fileContent string
		expected    []byte
		expectedErr error
	}{
		{
			name:        "Successful read",
			rootDir:     "/tmp",
			fileName:    "readtest1.txt",
			fileContent: "Hello",
			expected:    []byte("Hello"),
			expectedErr: nil,
		},
		{
			name:        "Failed read",
			rootDir:     "/tmp",
			fileName:    "readtest2.txt",
			fileContent: "",
			expected:    []byte(nil),
			expectedErr: &fs.PathError{
				Op:   "open",
				Path: "/tmp/readtest2.txt",
				Err:  errors.New("no such file or directory"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			filePath := fmt.Sprintf("%s/%s", tc.rootDir, tc.fileName)
			if tc.fileContent != "" {
				os.WriteFile(filePath, []byte(tc.fileContent), 0644)
			}

			actual, err := ReadFile(tc.rootDir, tc.fileName)

			if err != nil {
				assert.IsType(t, &fs.PathError{}, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.expected, actual)
		})
	}

	t.Cleanup(func() {
		for _, tc := range testCases {
			os.Remove(fmt.Sprintf("%s/%s", tc.rootDir, tc.fileName))
		}
	})
}
