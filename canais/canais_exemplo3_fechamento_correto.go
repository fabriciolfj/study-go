package canais

import (
	"log"
	"time"
)

func Execute3() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send2(ch) //
	for {        //
		select {
		case m, ok := <-ch: //
			if !ok { //
				log.Println("Channel closed.")
				return
			}
			log.Println("Got message:", m)
		case <-timeout: //
			log.Println("Time out")
			return
		default: //
			log.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send2(ch chan bool) { //
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	log.Println("Sent and closed")
}
