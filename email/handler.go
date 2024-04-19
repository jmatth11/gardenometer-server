package email

import (
	"fmt"
	"net/smtp"
)

const (
  SmtpHost = "smtp.gmail.com"
  SmtpPort = "587"
)

type EmailClient struct {
  To []string
  email string
  client smtp.Auth
}

func NewEmailClient(email string, password string) *EmailClient {
  auth := smtp.PlainAuth("", email, password, SmtpHost)
  return &EmailClient{
    To: make([]string, 0),
    client: auth,
    email: email,
  }
}

func (ec *EmailClient) SendMail(message string) error {
  return smtp.SendMail(fmt.Sprintf("%s:%s", SmtpHost, SmtpPort), ec.client, ec.email, ec.To, []byte(message))
}
