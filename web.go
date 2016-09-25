package main

import (
	"encoding/json"
	"fmt"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"github.com/LogvinovLeon/UberChallenge/email-providers/mailgun"
	"github.com/LogvinovLeon/UberChallenge/email-providers/sendgrid"
	"log"
	"net/http"
	"os"
)

type SenderType func(*definitions.EmailSendPayload)  error

func sendEmailWithFallback(payload *definitions.EmailSendPayload, primarySender, secondarySender SenderType) error {
	return primarySender(payload)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		var payload definitions.EmailSendPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(payload)
		var e error
		if payload.PreferredProvider == "mailgun" {
			e = sendEmailWithFallback(&payload, mailgun.Send, sendgrid.Send)
		} else {
			e = sendEmailWithFallback(&payload, sendgrid.Send, mailgun.Send)
		}
		if e != nil {
			log.Println(e)
			return
		}
	}
}

func main() {
	http.HandleFunc("/email/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
