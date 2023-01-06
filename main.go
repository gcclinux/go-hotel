package main

import (
	"log"
	"net/http"
)

const portNumber = ":4000"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting Application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}
