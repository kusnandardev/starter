package svcmail

import (
	"kusnandartoni/starter/pkg/mail"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/redisdb"

	"fmt"
	"io/ioutil"
)

// Verify :
type Verify struct {
	Email      string
	UserName   string
	VerifyLink string
}

// Store :
func (v *Verify) Store() error {
	return redisdb.StoreVerify(v)
}

// Send :
func (v *Verify) Send() {
	h := mail.Hermes{Product: getDefaultProduct()}
	h.Theme = new(mail.Default)
	generateEmails(h, getVerifyBody(v), "Verify")

	htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("runtime/mail/%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), "Verify"))
	if err != nil {
		panic(err)
	}
	txtBytes, err := ioutil.ReadFile(fmt.Sprintf("runtime/mail/%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), "Verify"))
	if err != nil {
		panic(err)
	}
	err = util.SendMail(v.Email, "Verify your email", string(htmlBytes), string(txtBytes))
	if err != nil {
		panic(err)
	}
}

func getVerifyBody(v *Verify) mail.Email {
	return mail.Email{
		Body: mail.Body{
			Name: v.UserName,
			Intros: []string{
				"Welcome to starter! Thank you for registering.",
			},
			Dictionary: []mail.Entry{
				// {Key: "Company Name", Value: s.CompanyName},
				// {Key: "Business Location", Value: s.BusinessLocation},
				// {Key: "Business category", Value: s.BusinessCategory},
			},
			Actions: []mail.Action{
				{
					Instructions: "Please Verify your email address by click this button:",
					Button: mail.Button{
						Color:     "#FFA500",
						TextColor: "#FFFFFF",
						Text:      "Verify",
						Link:      v.VerifyLink,
					},
				},
			},
			Outros: []string{
				"",
			},
			Signature: "Warm Regards,",
		},
	}
}
