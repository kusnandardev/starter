package svcmail

import (
	"kusnandartoni/starter/pkg/mail"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/redisdb"
	"fmt"
	"io/ioutil"
)

// ForgotPassword :
type ForgotPassword struct {
	Email     string
	UserName  string
	ResetLink string
}

// Store :
func (f *ForgotPassword) Store() error {
	return redisdb.StoreForgot(f)
}

// Send :
func (f *ForgotPassword) Send() {
	h := mail.Hermes{Product: getDefaultProduct()}
	h.Theme = new(mail.Default)
	generateEmails(h, getForgetBody(f), "Forgot")
	htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("runtime/mail/%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), "Forgot"))
	if err != nil {
		panic(err)
	}
	txtBytes, err := ioutil.ReadFile(fmt.Sprintf("runtime/mail/%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), "Forgot"))
	if err != nil {
		panic(err)
	}
	err = util.SendMail(f.Email, "Request Reset Password", string(htmlBytes), string(txtBytes))
	if err != nil {
		panic(err)
	}
}

func getForgetBody(f *ForgotPassword) mail.Email {
	return mail.Email{
		Body: mail.Body{
			Name: f.UserName,
			Intros: []string{
				"You have received this email because a password reset request for starter account was received.",
			},
			Actions: []mail.Action{
				{
					Instructions: "Click the button below to reset your password:",
					Button: mail.Button{
						Color:     "#FFA500",
						TextColor: "#FFFFFF",
						Text:      "Reset your password",
						Link:      f.ResetLink,
					},
				},
			},
			Outros: []string{
				"If you did not request a password reset, no further action is required on your part.",
			},
			Signature: "Warm Regards,",
		},
	}
}
