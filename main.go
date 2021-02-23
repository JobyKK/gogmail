package gogmail

import (
	"fmt"
	"net/smtp"
)

const (
	emailHost = "smtp.gmail.com"
	emailPort = 587
)

func NewGmailClient(email, password string) *GmailClient {
	auth := smtp.PlainAuth("", email, password, emailHost)
	server := fmt.Sprintf("%s:%v", emailHost, emailPort)

	return &GmailClient{
		emailFrom:  email,
		smtpAuth:   auth,
		smtpServer: server,
	}
}

func (client GmailClient) SendEmail(to []string, data *GmailSendData) (bool, error) {
	subject := data.Subject
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + data.Body)

	err := smtp.SendMail(client.smtpServer, client.smtpAuth, client.emailFrom, to, msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
