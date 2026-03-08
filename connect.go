package main

import (
	"fmt"
	"net"
	"time"
)

func connect(host, port string) (net.Conn, error) {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", address, err)
	}
	return conn, nil
}
