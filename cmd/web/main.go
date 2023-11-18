package main

import (
	"encoding/gob"
	"log"
	"myapp/internal/config"
	"myapp/internal/driver"
	"myapp/internal/handlers"
	"myapp/internal/helpers"
	"myapp/internal/models"
	"myapp/internal/render"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	from := "me@here.com"
	to := "you@here.com"
	auth := smtp.PlainAuth("", from, "", "localhost")
	err = smtp.SendMail("localhost:1025", auth, from, []string{to}, []byte("Hello from go-hotel"))
	if err != nil {
		log.Println("mail in main.go: ", err)
	}

	log.Println("Starting Application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// Store data that we will be keeping in the session.
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// Change this to true if in prodution
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false // false = cooky goes when browser closes
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connet to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=odroid port=5432 dbname=hotel user=testie password=pastie")
	if err != nil {
		log.Fatal("main.go cannot connect to database -->", err)
		return nil, err
	}
	log.Println("Connected to database...")

	myTemplateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("main.go cannot create template cache -->", err)
		return nil, err
	}

	app.TemplateCache = myTemplateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
