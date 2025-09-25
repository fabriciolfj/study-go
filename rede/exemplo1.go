package rede

import (
	"log"
	"net"
)

func Execute() {
	conn, err := net.Dial("tcp", "127.0.0.1:1902")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example", f)

	logger.Println("Hello World")
	logger.Panicf("this is panic")
}
