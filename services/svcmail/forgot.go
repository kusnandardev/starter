package svcmail

import (
	"fmt"
	"io/ioutil"
	"strings"

	"kusnandartoni/starter/pkg/mail"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/redisdb"
)

// Forgot :
type Forgot struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	ButtonLink string `json:"button_link"`
}

// Store :
func (f *Forgot) Store() error {
	return redisdb.StoreForgot(f)
}

// Send :
func (f *Forgot) Send() error {
	subjectEmail := "Permintaan Pergantian Password"
	fileName := fmt.Sprintf("Forgot-%s", strings.ReplaceAll(f.Email, "@", "."))
	filePath := fmt.Sprintf("%s/%s", "runtime/mail", fileName)

	h := mail.Mail{MailType: "forgot"}
	generateEmailTemplate(h, getForgotBody(f), fileName)

	htmlBytes, err := ioutil.ReadFile(filePath + ".html")
	if err != nil {
		return err
	}
	txtBytes, err := ioutil.ReadFile(filePath + ".txt")
	if err != nil {
		return err
	}
	err = util.SendMail(f.Email, subjectEmail, string(htmlBytes), string(txtBytes))
	if err != nil {
		return err
	}
	return nil
}

func getForgotBody(f *Forgot) mail.Format {
	return mail.Format{
		Forgot: mail.Forgot{
			Name:       f.Name,
			ButtonLink: f.ButtonLink,
		},
	}
}
