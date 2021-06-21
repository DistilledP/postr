package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPServer(t *testing.T) {
	grpcServer := NewGRPCServer()
	httpServer := NewHTTPServer(grpcServer)

	assert.IsType(t, &http.Server{}, httpServer)
}
