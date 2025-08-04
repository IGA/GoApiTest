package main

import (
	"go-api/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.PingHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
