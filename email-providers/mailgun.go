package mailgun

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"os"
	"github.com/LogvinovLeon/UberChallenge/definitions"
)

type MailgunSender struct {
}

func (mailgunSender MailgunSender) Send(messagePayload *definitions.EmailSendPayload) (sendResult definitions.SendResult, err error) {
	domain := "uberchallenge.email"
	mg := mailgun.NewMailgun(domain, os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_KEY"))
	message := mailgun.NewMessage("admin@uberchallenge.email",
		messagePayload.Subject,
		messagePayload.Body,
		messagePayload.To)
	response, id, err := mg.Send(message)
	fmt.Println(response, id, err)
	sendResult = definitions.SendResult{Provider: "mailgun", Id: id}
	return
}
