package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawURL, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	host, port, path, err := parseURL(rawURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connecting to", host)

	var b strings.Builder
	fmt.Fprintf(&b, "GET %s HTTP/1.1\r\n", path)
	fmt.Fprintf(&b, "Host: %s\r\n", host)
	fmt.Fprintf(&b, "Accept: */*\r\n")
	fmt.Fprintf(&b, "Connection: close\r\n")
	fmt.Fprintf(&b, "\r\n")

	fmt.Print(b.String())

	conn, err := connect(host, port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := sendRequest(conn, b.String()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	response, err := readResponse(conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print(string(response))
}
