package main

import (
	"flag"
	"fmt"
)

func parseArgs(args []string) (string, error) {
	// set the flag set command for curl
	fs := flag.NewFlagSet("curl", flag.ContinueOnError)

	if err := fs.Parse(args); err != nil {
		return "", err
	}

	// check arguments are 1 - the URL
	rest := fs.Args()
	if len(rest) != 1 {
		return "", fmt.Errorf("usage: cccurl <url>")
	}

	// extract url and connect
	return rest[0], nil
}
