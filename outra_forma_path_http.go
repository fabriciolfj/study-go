package main

import (
	"fmt"
	"net/http"
)

func test9() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler3)
	mux.HandleFunc("GET /goodbye/", goodbyeHandler3)
	mux.HandleFunc("GET /goodbye/{name}", goodbyeHandler3)

	if err := http.ListenAndServe(":8084", mux); err != nil {
		panic(err)
	}
}

func helloHandler3(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbyeHandler3(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
