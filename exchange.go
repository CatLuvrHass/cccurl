package main

import (
	"fmt"
	"io"
	"net"
)

func sendRequest(conn net.Conn, request string) error {
	_, err := conn.Write([]byte(request))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	return nil
}

func readResponse(conn net.Conn) ([]byte, error) {
	response, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	return response, nil
}