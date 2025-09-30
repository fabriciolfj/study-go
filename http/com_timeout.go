package http

import (
	"fmt"
	"net/http"
	"time"
)

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("you should never see me"))
}

func Execute2() {
	muxer := http.NewServeMux()
	muxer.HandleFunc("/timeout", timeoutHandler)

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		Handler:      muxer,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}
}
