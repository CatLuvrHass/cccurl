package main

import (
	"fmt"
	"net"
	"net/url"
)

func parseURL(rawURL string) (host string, port string, path string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid url")
	}

	if u.Scheme == "" {
		return "", "", "", fmt.Errorf("scheme missing (use http://)")
	}
	if u.Scheme != "http" {
		return "", "", "", fmt.Errorf("only http is supported")
	}

	if u.Host == "" {
		return "", "", "", fmt.Errorf("host missing")
	}
	// host / port
	h, p, splitErr := net.SplitHostPort(u.Host)
	if splitErr == nil {
		host = h
		port = p
	} else {
		host = u.Host
		port = "80"
	}

	// path + query
	path = u.Path
	if path == "" {
		path = "/"
	}
	if u.RawQuery != "" {
		path = path + "?" + u.RawQuery
	}

	return host, port, path, nil

}
