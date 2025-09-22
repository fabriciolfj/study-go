package canais

import (
	"fmt"
	"time"
)

func Execute2() {
	msg := make(chan string)
	until := time.After(5 * time.Second)
	go send(msg)
	for {
		select {
		case msg := <-msg:
			fmt.Println(msg)
		case <-until:
			close(msg)
			time.Sleep(500 * time.Millisecond)
			break
		}
	}

}

func send(ch chan string) {
	for {
		ch <- "Hello"
		time.Sleep(500 * time.Millisecond)
	}
}
