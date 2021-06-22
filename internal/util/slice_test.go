package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlicePop(t *testing.T) {
	tests := []struct {
		name              string
		items             []string
		expected          string
		expectedRemainder []string
	}{
		{
			name:              "Pops last item",
			items:             []string{"item1", "item2", "item3"},
			expected:          "item3",
			expectedRemainder: []string{"item1", "item2"},
		},
		{
			name:              "Pops only item",
			items:             []string{"item1"},
			expected:          "item1",
			expectedRemainder: []string{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, actualRemainder := StringSlicePop(tc.items)

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedRemainder, actualRemainder)
		})
	}
}
