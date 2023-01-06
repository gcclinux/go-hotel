package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":4000"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// renderTemplate is where we will be rendering all templates through
func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles("./templates/" + page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("error parsing templates:", err)
		return
	}
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting Application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}
