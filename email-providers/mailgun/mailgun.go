package mailgun

import (
	"github.com/mailgun/mailgun-go"
	"os"
	"github.com/LogvinovLeon/UberChallenge/definitions"
)

func Send(messagePayload *definitions.EmailSendPayload) error {
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_KEY"))
	message := mailgun.NewMessage("admin@uberchallenge.email",
		messagePayload.Subject,
		messagePayload.Body,
		messagePayload.To)
	_, _, err := mg.Send(message)
	return err
}
