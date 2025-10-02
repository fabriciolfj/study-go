package client_http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Execute2() {
	cc := &http.Client{Timeout: time.Second * 5}
	res, err := cc.Get("http://www.manning.com")
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s", b)
}
