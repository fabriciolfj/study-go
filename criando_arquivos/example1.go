package criando_arquivos

import "os"

func Execute() {
	file, err := os.Create("teste.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("Hello World")
}
