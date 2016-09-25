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

func sendEmailWithFallback(payload *definitions.EmailSendPayload, primarySender, secondarySender definitions.EmailSender) (definitions.SendResult, error) {
	return primarySender.Send(payload)
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
		var r definitions.SendResult
		var e error
		if payload.PreferredProvider == "mailgun" {
			r, e = sendEmailWithFallback(&payload, mailgun.MailgunSender{}, sendgrid.SendGridSender{})
		} else {
			r, e = sendEmailWithFallback(&payload, sendgrid.SendGridSender{}, mailgun.MailgunSender{})
		}
		if e != nil {
			log.Println(e)
			return
		}
		log.Println(r)
	}
}

func main() {
	http.HandleFunc("/email/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
