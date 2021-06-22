package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImageDir(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		envVal   string
	}{
		{
			name:     "No env set",
			expected: DefaultImageDir,
		},
		{
			name:     "Env set",
			expected: "./",
			envVal:   "./",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv(EnvImageDirKey)
			if tc.envVal != "" {
				os.Setenv(EnvImageDirKey, tc.envVal)
			}

			actual := GetImageDir()

			assert.Equal(t, tc.expected, actual)
		})
	}

	t.Cleanup(func() {
		os.Unsetenv(EnvImageDirKey)
		os.Remove("config_test_dir")
	})
}

func TestGetAddress(t *testing.T) {
	testCases := []struct {
		name       string
		defaultVal string
		expected   string
		key        string
		envVal     string
	}{
		{
			name:       "Returns default value",
			defaultVal: "3000",
			expected:   ":3000",
			key:        "TEST_KEY_1",
		},
		{
			name:       "Returns environment value",
			defaultVal: "3000",
			expected:   ":3002",
			key:        "TEST_KEY_2",
			envVal:     "3002",
		},
		{
			name:       "Handles hostname in address",
			defaultVal: "3000",
			expected:   "localhost:3002",
			key:        "TEST_KEY_3",
			envVal:     "localhost:3002",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envVal != "" {
				os.Setenv(tc.key, tc.envVal)
			}

			assert.Equal(t, GetAddress(tc.key, tc.defaultVal), tc.expected)

			if os.Getenv(tc.key) != "" {
				os.Unsetenv(tc.key)
			}
		})
	}
}
