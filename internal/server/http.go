package server

import (
	"net/http"

	"github.com/DistilledP/postr/internal/server/handler"
	"google.golang.org/grpc"
)

const DefaultHTTPPort = "3000"

func NewHTTPServer(grpcServer *grpc.Server) *http.Server {
	httpRouter := http.NewServeMux()
	httpRouter.Handle("/", &handler.RootHandler{})

	s := &http.Server{
		Handler: httpRouter,
	}

	s.RegisterOnShutdown(func() {
		grpcServer.GracefulStop()
	})

	return s
}
