package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func GetAddress(key, defaultValue string) string {
	address := defaultValue
	if envAddress := strings.TrimSpace(os.Getenv(key)); envAddress != "" {
		address = envAddress
	}

	if !strings.Contains(address, ":") {
		address = ":" + address
	}

	return address
}

func OpenConn(addr, proto string) (net.Listener, error) {
	if proto == "" {
		proto = "tcp"
	}

	sock, err := net.Listen(proto, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s for protocol %s: %v", addr, proto, err)
	}

	return sock, nil
}
