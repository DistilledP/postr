package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	pb "github.com/DistilledP/postr/internal/proto"
)

type stubHandler struct {
	mock.Mock
}

func (m *stubHandler) handler(ctx context.Context, req interface{}) (interface{}, error) {
	args := m.Called(ctx, req)

	return args.Get(0), args.Error(1)
}

func TestMimeMismatchInterceptor(t *testing.T) {
	testCases := []struct {
		name             string
		req              interface{}
		expectedResponse interface{}
		handlerCalled    bool
	}{
		{
			"Invalid type",
			&pb.ImageUpload{
				Payload:  []byte("BMP***"),
				FileType: pb.MimeType_JPEG,
			},
			&pb.ImageUploadResponse{
				ErrorMessage: `File format mis-match, expecting "JPEG", got "UNKNOWN"`,
			},
			false,
		},
		{
			"Valid type",
			&pb.ImageUpload{
				Payload:  []byte("\xFF\xD8\xFF"),
				FileType: pb.MimeType_JPEG,
			},
			&pb.ImageUploadResponse{
				Status: pb.Status_SUCCESS,
			},
			true,
		},
		{
			"Ignored request type",
			nil,
			&pb.ImageUploadResponse{
				Status: pb.Status_FAILED,
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			serverInfo := &grpc.UnaryServerInfo{}

			mockHandler := new(stubHandler)
			if tc.handlerCalled {
				mockHandler.On("handler", ctx, tc.req).Return(tc.expectedResponse, nil)
			}

			resp, err := mimeMismatchInterceptor(ctx, tc.req, serverInfo, mockHandler.handler)

			assert.Nil(t, err)

			assert.IsType(t, tc.expectedResponse, resp)
			assert.Equal(t, tc.expectedResponse, resp)

			if tc.handlerCalled {
				mockHandler.AssertExpectations(t)
			} else {
				mockHandler.AssertNotCalled(t, "handler")
			}
		})
	}
}
