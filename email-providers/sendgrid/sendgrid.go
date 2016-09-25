package sendgrid

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/LogvinovLeon/UberChallenge/definitions"
	"os"
)

type SendGridSender struct {
}

func (sendGridSender SendGridSender) Send(message *definitions.EmailSendPayload) (sendResult definitions.SendResult, err error) {
	from := mail.NewEmail("UberChallenge", "admin@uberchallenge.email")
	subject := message.Subject
	to := mail.NewEmail("Leonid Logvinov", message.To)
	content := mail.NewContent("text/plain", message.Body)
	m := mail.NewV3MailInit(from, subject, to, content)
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	sendResult = definitions.SendResult{Provider:"sendgrig", Id: response.Body}
	return
}