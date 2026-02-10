package main

import (
	"fmt"
	"os"
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
	fmt.Printf("Sending request GET %s HTTP/1.1\n", path)
	fmt.Printf("Host: %s\n", host)
	fmt.Println("Accept: */*")

	_ = port // used in Step 2
}
