package main

import (
	"encoding/json"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"github.com/LogvinovLeon/UberChallenge/email-providers/mailgun"
	"github.com/LogvinovLeon/UberChallenge/email-providers/sendgrid"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"log"
)

type SenderType func(*definitions.EmailSendPayload) error

func sendEmailWithFallback(payload *definitions.EmailSendPayload, primarySender, secondarySender SenderType) error {
	log.Println("Trying to send using prefered provider...")
	err := primarySender(payload)
	if err != nil {
		log.Println("Failed. Try with the fallback provider...")
		err = secondarySender(payload)
		return err
	} else {
		return err
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodPost {
		if r.Body == nil {
			log.Println("Empty body")
			http.Error(w, "Please send a request body", 400)
			return
		}
		var payload definitions.EmailSendPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			log.Println("Not a correct JSON")
			http.Error(w, err.Error(), 400)
			return
		}
		var e error
		if payload.PreferredProvider == "sendgrid" {
			e = sendEmailWithFallback(&payload, sendgrid.Send, mailgun.Send)
		} else {
			// Mailgun is a default option
			e = sendEmailWithFallback(&payload, mailgun.Send, sendgrid.Send)
		}
		if e != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write([]byte("OK!\n"))
	}
}

func main() {
	http.HandleFunc("/email", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
