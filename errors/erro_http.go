package errors

import (
	"errors"
	"net/http"
)

func ExecuteHttp() {
	http.HandleFunc("GET /", handler)
	/*defer func() { //dentro do http ja tem para nao parar a app de funcionar por causa do panic
		if err := recover(); err != nil {
			println(err)
		}
	}()*/

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("could not start server_grpc")
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	panic(errors.New("Fake panic!"))
}
