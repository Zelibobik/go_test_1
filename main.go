package main

import (
	"log"
	"net/http"
	"github.com/zelibobik/go_test_1/handlers"
)

// How to try it: go run main.go
func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}