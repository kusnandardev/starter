package svcmail

import (
	"kusnandartoni/starter/pkg/mail"
	"fmt"
	"io/ioutil"
	"os"
)

func generateEmails(h mail.Hermes, email mail.Email, templateName string) {
	// Generate the HTML template and save it
	res, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll("runtime/mail/"+h.Theme.Name(), 0744)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("runtime/mail/%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), templateName), []byte(res), 0644)
	if err != nil {
		panic(err)
	}

	// Generate the plaintext template and save it
	res, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("runtime/mail/%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), templateName), []byte(res), 0644)
	if err != nil {
		panic(err)
	}
}

func getDefaultProduct() mail.Product {
	return mail.Product{
		Name:        "starter Team",
		Link:        "https://saviory.tech",
		Logo:        "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		TroubleText: "Or copy and paste the following link into your browser",
		Copyright:   "Copyright © 2019 starter All rights reserved.",
	}
}
