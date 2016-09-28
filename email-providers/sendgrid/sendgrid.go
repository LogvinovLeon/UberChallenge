package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"os"
	"errors"
	"log"
)

func Send(message *definitions.EmailSendPayload) error {
	log.Println("Sending with sendgrid")
	from := mail.NewEmail("", "admin@uberchallenge.email")
	subject := message.Subject
	to := mail.NewEmail("", message.To)
	content := mail.NewContent("text/plain", message.Body)
	m := mail.NewV3MailInit(from, subject, to, content)
	SENDGRID_API_KEY, set := os.LookupEnv("SENDGRID_API_KEY")
	if !set {
		return errors.New("SENDGRID_API_KEY env variable is not set")
	}
	request := sendgrid.GetRequest(SENDGRID_API_KEY, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)
	return err
}