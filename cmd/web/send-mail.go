package main

import (
	"log"
	"myapp/internal/models"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	// Example to listen to a channel but not used here
	// m := <-app.MailChan

	// I'm going to fire off in the background, start a go routine by using that command go and I'm going
	// to fire an unnamed function or an anonymous function.
	// So to do that I just say func and then open and close parentheses and I'm not putting any parameters

	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = mailPort
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.Password = "Passw0rd"

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	//email.SetBody(mail.TextHTML, "Hello, <strong>World!</strong>")
	email.SetBody(mail.TextHTML, string(m.Content))

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent!")
	}
}
