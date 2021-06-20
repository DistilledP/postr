package middleware

import (
	"testing"

	"github.com/DistilledP/postr/internal/test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUnaryServerMiddleware(t *testing.T) {
	expectedOrder := []grpc.UnaryServerInterceptor{
		mimeTypeCheckInterceptor,
		mimeMismatchInterceptor,
	}
	assert.Len(t, UnaryServerMiddleware, len(expectedOrder))

	for i, actual := range UnaryServerMiddleware {
		expectedFunctionName := test.GetFunctionName(expectedOrder[i])
		actualFunctionName := test.GetFunctionName(actual)

		assert.Equal(t, expectedFunctionName, actualFunctionName)
	}
}
