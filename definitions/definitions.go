package definitions

type EmailSendPayload struct {
	To      string
	Subject string
	Body    string
	PreferredProvider string
}