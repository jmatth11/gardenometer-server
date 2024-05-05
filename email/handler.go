package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

const (
  SmtpHost = "smtp.gmail.com"
  SmtpPort = "587"
)

type EmailClient struct {
  To []string
  email string
  password string
}

func NewEmailClient(email string, password string) *EmailClient {
  return &EmailClient{
    To: make([]string, 0),
    email: email,
    password: password,
  }
}

func (ec *EmailClient) SendMail(message string) error {
  auth := smtp.PlainAuth("", ec.email, ec.password, SmtpHost)
  bodyMessage := bytes.Buffer{}
  bodyMessage.WriteString("From: ")
  bodyMessage.WriteString(ec.email)
  bodyMessage.WriteString("\\r\\n")
  bodyMessage.WriteString("To: ")
  bodyMessage.WriteString(strings.Join(ec.To, ", "))
  bodyMessage.WriteString("\\r\\n")
  bodyMessage.WriteString("Subject: Gardenometer Update\\r\\n")
  bodyMessage.WriteString(message)
  return smtp.SendMail(fmt.Sprintf("%s:%s", SmtpHost, SmtpPort), auth, ec.email, ec.To, bodyMessage.Bytes())
}
