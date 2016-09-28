package definitions

type EmailSendPayload struct {
	To      string `valid:"email" json:"to"`
	Subject string `valid:"-" json:"subject"`
	Body    string `valid:"-" json:"body"`
	PreferredProvider string `json:"preferred_provider"`
}