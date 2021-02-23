package gogmail

import "net/smtp"

type GmailSendData struct {
	Subject string
	Body    string
}

type GmailClient struct {
	emailFrom  string
	smtpAuth   smtp.Auth
	smtpServer string
}
