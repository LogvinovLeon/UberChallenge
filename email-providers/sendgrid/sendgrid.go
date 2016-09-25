package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"os"
)

func Send(message *definitions.EmailSendPayload) error {
	from := mail.NewEmail("UberChallenge", "admin@uberchallenge.email")
	subject := message.Subject
	to := mail.NewEmail("Leonid Logvinov", message.To)
	content := mail.NewContent("text/plain", message.Body)
	m := mail.NewV3MailInit(from, subject, to, content)
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)
	return err
}