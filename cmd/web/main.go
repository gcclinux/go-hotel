package main

import (
	"encoding/gob"
	"log"
	"myapp/internal/config"
	"myapp/internal/handlers"
	"myapp/internal/models"
	"myapp/internal/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting Application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// Store data that we will be keeping in the session.
	gob.Register(models.Reservation{})

	// Change this to true if in prodution
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false // false = cooky goes when browser closes
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	myTemplateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("main.go cannot create template cache -->", err)
		return err
	}

	app.TemplateCache = myTemplateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil
}
