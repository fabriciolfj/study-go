package rede

import (
	"log"
	"net"
	"time"
)

func Execute2() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("tcp", "127.0.0.1:1902", timeout)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example", f)

	logger.Println("Hello World")
	logger.Panicf("this is panic")
}
