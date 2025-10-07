package ips

import (
	"log"
	"net"
	"os"
)

func Execute() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		log.Println(addr)
	}
}
