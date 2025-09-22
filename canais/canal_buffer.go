package canais

import (
	"log"
	"time"
)

func Execute5() {
	lock := make(chan bool, 1) //
	for i := 1; i < 7; i++ {   //
		go worker(i, lock) //
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	log.Printf("%d wants the lock\n", id)
	lock <- true                        //
	log.Printf("%d has the lock\n", id) //
	<-lock                              //
	log.Printf("%d is releasing the lock\n", id)
}
