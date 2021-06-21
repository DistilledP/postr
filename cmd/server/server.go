package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/DistilledP/postr/internal/server"
	"github.com/DistilledP/postr/internal/util"
)

type fn func(...interface{})

func init() {
	util.GetImageDir() // Ensure the directory is created
}

func main() {
	grpcAddr := util.GetAddress("GRPC_PORT", server.DefaultGRPCPort)
	httpAddr := util.GetAddress("HTTP_PORT", server.DefaultHTTPPort)

	grpcServer := server.NewGRPCServer()
	httpServer := server.NewHTTPServer(grpcServer)

	grpcSocket, err := server.OpenConn(grpcAddr, "tcp")
	if err != nil {
		log.Fatal(err)
	}
	defer grpcSocket.Close()
	log.Printf("GRPC server listening on tcp port %s", strings.TrimLeft(grpcAddr, ":"))

	httpSocket, err := server.OpenConn(httpAddr, "tcp")
	if err != nil {
		log.Fatal(err)
	}
	defer httpSocket.Close()
	log.Printf("HTTP server listening on tcp port %s", strings.TrimLeft(httpAddr, ":"))

	go func(logger fn) {
		if err := grpcServer.Serve(grpcSocket); err != nil {
			logger(fmt.Errorf("failed to start grpc server: %v", err))
		}
	}(log.Fatal)

	if err = httpServer.Serve(httpSocket); err != nil {
		log.Fatal(fmt.Errorf("failed to start http server: %v", err))
	}
}
