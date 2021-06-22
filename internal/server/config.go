package server

import (
	"fmt"
	"net"
)

func OpenConn(addr, proto string) (net.Listener, error) {
	if proto == "" {
		proto = "tcp"
	}

	sock, err := net.Listen(proto, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s for protocol %s", addr, proto)
	}

	return sock, nil
}
