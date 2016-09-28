package mailgun

import (
	"github.com/mailgun/mailgun-go"
	"os"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"errors"
	"log"
)

func Send(messagePayload *definitions.EmailSendPayload) error {
	log.Println("Sending with mailgun")
	MAILGUN_DOMAIN, set := os.LookupEnv("MAILGUN_DOMAIN")
	if !set {
		return errors.New("MAILGUN_DOMAIN env variable is not set")
	}
	MAILGUN_API_KEY, set := os.LookupEnv("MAILGUN_API_KEY")
	if !set {
		return errors.New("MAILGUN_API_KEY env variable is not set")
	}
	MAILGUN_PUBLIC_KEY, set := os.LookupEnv("MAILGUN_PUBLIC_KEY")
	if !set {
		return errors.New("MAILGUN_PUBLIC_KEY env variable is not set")
	}
	mg := mailgun.NewMailgun(MAILGUN_DOMAIN, MAILGUN_API_KEY, MAILGUN_PUBLIC_KEY)
	message := mailgun.NewMessage("admin@uberchallenge.email",
		messagePayload.Subject,
		messagePayload.Body,
		messagePayload.To)
	_, _, err := mg.Send(message)
	return err
}
