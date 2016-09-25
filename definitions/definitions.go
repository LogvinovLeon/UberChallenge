package definitions

type EmailSendPayload struct {
	To      string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
	PreferredProvider string
}

type SendResult struct {
	Provider string
	Id string
}

type EmailSender interface {
	Send(messagePayload *EmailSendPayload) (SendResult, error)
}
