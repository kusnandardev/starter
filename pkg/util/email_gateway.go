package util

import (
	"kusnandartoni/starter/pkg/setting"
	"net/mail"

	"github.com/go-gomail/gomail"
)

// SendMail :
func SendMail(
	to string,
	subject string,
	htmlBody string,
	txtBody string) error {

	smtp := setting.SMTPSetting

	from := mail.Address{
		Name:    smtp.SMTPIdentity,
		Address: smtp.SMTPSenderEmail,
	}

	m := gomail.NewMessage()
	m.Reset()
	m.SetHeader("From", from.String())
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtp.SMTPServer, smtp.SMTPPort, smtp.SMTPUser, smtp.SMTPPasswd)

	return d.DialAndSend(m)
}
