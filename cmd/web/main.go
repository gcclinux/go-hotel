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
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
const mailPort = 1025

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

	// Single line required to close the channel as soon as connection request no longer required
	defer db.SQL.Close()

	// Single line required to close the channel as soon as message sent rather than leaving it open forever
	defer close(app.MailChan)

	// Start listening again for the next meail
	listenForMail()

	// Test message when server started
	// msg := models.MailData{
	// 	To:      "you@there.com",
	// 	From:    "me@here.com",
	// 	Subject: "Server Started",
	// 	Content: "Server Started from <strong><i>command line!</i></strong>",
	// }

	// app.MailChan <- msg

	log.Println("Starting Application on port", portNumber)
	log.Print("Starting Mail listen on port :", mailPort, "\n")

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

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

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
