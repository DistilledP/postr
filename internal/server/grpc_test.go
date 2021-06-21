package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestNewGRPCServer(t *testing.T) {
	expectedServices := map[string]grpc.ServiceInfo{
		"Postr": {
			Methods: []grpc.MethodInfo{
				{
					Name:           "Upload",
					IsClientStream: false,
					IsServerStream: false,
				},
			},
		},
	}

	grpcServer := NewGRPCServer()
	serviceInfo := grpcServer.GetServiceInfo()

	var keys []string
	for k := range serviceInfo {
		keys = append(keys, k)
	}

	for expectedSvc, expectedType := range expectedServices {
		assert.Contains(t, keys, expectedSvc)
		assert.Equal(t, expectedType.Methods, serviceInfo[expectedSvc].Methods)
	}
}
