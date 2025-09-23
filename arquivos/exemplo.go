package arquivos

import (
	"log"
	"os"
)

func Execute() {
	data, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	log.Println(string(data))
}
