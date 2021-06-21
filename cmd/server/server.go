package main

import (
	"fmt"
	"log"

	"github.com/DistilledP/postr/internal/server"
)

type fn func(...interface{})

func main() {
	grpcAddr := server.GetAddress("GRPC_PORT", server.DefaultGRPCPort)
	httpAddr := server.GetAddress("HTTP_PORT", server.DefaultHTTPPort)

	grpcServer := server.NewGRPCServer()
	httpServer := server.NewHTTPServer(grpcServer)

	grpcSocket, err := server.OpenConn(grpcAddr, "tcp")
	if err != nil {
		log.Fatal(err)
	}
	defer grpcSocket.Close()

	httpSocket, err := server.OpenConn(httpAddr, "tcp")
	if err != nil {
		log.Fatal(err)
	}
	defer httpSocket.Close()

	go func(logger fn) {
		if err := grpcServer.Serve(grpcSocket); err != nil {
			logger(fmt.Errorf("failed to start grpc server: %v", err))
		}
	}(log.Fatal)

	if err = httpServer.Serve(httpSocket); err != nil {
		log.Fatal(fmt.Errorf("failed to start http server: %v", err))
	}
}
