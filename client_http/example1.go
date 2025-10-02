package client_http

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func Execute() {
	req, err := http.Get("https://example.com")
	if hasTimeOut(err) {
		panic("request has timeout")
	}

	if err != nil {
		panic("something else has happened")
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	fmt.Println("%s", b)
}

func hasTimeOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}

	return false
}
