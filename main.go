package main

import (
	"log"
	"net/http"
)

const portNumber = ":4000"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {

}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting Application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}
