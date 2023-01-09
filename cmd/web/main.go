package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":5000"

// main is the main application function
func main() {

	var app config.AppConfig

	myTemplateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create templatr cache", err)
	}

	app.TemplateCache = myTemplateCache

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting Application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}
