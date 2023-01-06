package main

import (
	"log"
	"myapp/pkg/handlers"
	"net/http"
)

const portNumber = ":4000"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting Application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}
