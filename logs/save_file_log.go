package logs

import (
	"log"
	"os"
)

func Execute2() {
	file, err := os.OpenFile("logging.log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic("could not open file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LUTC | log.Lshortfile)
	log.Printf("display this message")
}
