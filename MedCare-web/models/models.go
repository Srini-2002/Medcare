package models

type MailGunEmail struct {
	RecipientEmail string `json:"recipientemail"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}
