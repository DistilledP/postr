package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileName(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "File name test 1",
			filePath: "/tmp/foo/bar/moo.txt",
			expected: "moo.txt",
		},
		{
			name:     "File name test 2",
			filePath: "/boo.txt",
			expected: "boo.txt",
		},
		{
			name:     "File name test 3",
			filePath: "shoo.txt",
			expected: "shoo.txt",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FileName(tc.filePath)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
