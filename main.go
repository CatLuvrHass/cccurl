package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {

	// set the flag set command for curl
	fs := flag.NewFlagSet("curl", flag.ContinueOnError)
	err := fs.Parse(os.Args[1:])

	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	// check arguments are 1 - the URL
	args := fs.Args()
	if len(args) != 1 {
		fmt.Println("Unexpected input")
		return
	}

	// extract url and connect
	rawURL := args[0]
	fmt.Println("connecting to", rawURL)

	// handle parsed URL
	u, err := url.Parse(rawURL)

	// if parsed url is no good, throw an error
	if err != nil {
		fmt.Println("error fatal")
		return
	}

	if u.Scheme == "" {
		fmt.Println("user should include http://")
		return
	}
	if u.Scheme != "http" {
		fmt.Println("only http is supported")
		return
	}

	fmt.Println(u.Host)
	if u.Host == "" {
		fmt.Println("host is missing")
		return
	}

	ph := strings.Split(u.Host, ":")
	var host string
	var port string
	if len(ph) >= 2 {
		host = ph[0]
		port = ph[1]
	} else {
		host = u.Host
		port = "80"
	}

	fmt.Println(host, port)

	path := u.Path
	if path == "" {
		path = "/"
	}

	if u.Path != "" {
		path = path + "?" + u.RawQuery
	}

	fmt.Println(path)
}
