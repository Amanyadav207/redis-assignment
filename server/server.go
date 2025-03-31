package server

import (
	"fmt"
	"key-value-cache/config"
	"key-value-cache/handlers"
	"net/http"
)

func Start() {
	http.HandleFunc("/put", handlers.PutHandler)
	http.HandleFunc("/get", handlers.GetHandler)
	port := config.GetPort()
	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, nil)
}