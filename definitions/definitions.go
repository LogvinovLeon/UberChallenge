package definitions

type EmailSendPayload struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
	PreferredProvider string
}