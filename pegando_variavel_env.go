package main

import (
	"fmt"
	"net/http"
	"os"
)

func test5() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	fmt.Println("Servidor rodando na porta", port)
	http.HandleFunc("/", homePage2)
	http.ListenAndServe(":"+port, nil)
}

func homePage2(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}

	fmt.Fprintf(res, "Hello World!")
}
