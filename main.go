package main

import (
	"fmt"
	"go-websocket/services"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/ws", services.HandleConnections)

	go services.HandleMessages()

	fmt.Println("Application up and running")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
