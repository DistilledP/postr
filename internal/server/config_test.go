package server

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	testCases := []struct {
		defaultVal string
		expected   string
		key        string
		envVal     string
	}{
		{
			defaultVal: "3000",
			expected:   ":3000",
			key:        "TEST_KEY_1",
		},
		{
			defaultVal: "3000",
			expected:   ":3002",
			key:        "TEST_KEY_2",
			envVal:     "3002",
		},
		{
			defaultVal: "3000",
			expected:   "localhost:3002",
			key:        "TEST_KEY_3",
			envVal:     "localhost:3002",
		},
	}

	for _, tc := range testCases {
		if tc.envVal != "" {
			os.Setenv(tc.key, tc.envVal)
		}

		assert.Equal(t, GetAddress(tc.key, tc.defaultVal), tc.expected)

		if os.Getenv(tc.key) != "" {
			os.Unsetenv(tc.key)
		}
	}
}
