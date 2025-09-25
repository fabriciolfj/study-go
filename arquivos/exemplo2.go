package arquivos

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Execute2() {
	file, err := os.Open("structured.log")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("File: name is %s, mode is %v, size is %d, Is directory: %v", info.Name(), info.Mode(), info.Size(), info.IsDir()))

	lineJson := make(map[string]interface{})
	var bChunk []byte
	for {
		b := make([]byte, 639)
		_, err := file.Read(b)

		if err != nil {
			break
		}

		bChunk = append(bChunk, b[0:]...)
		if err := json.Unmarshal(bChunk, &lineJson); err == nil {
			log.Println(lineJson)
			bChunk = []byte{}
		} else {
			log.Println(err)
		}
	}

}
