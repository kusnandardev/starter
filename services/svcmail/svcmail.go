package svcmail

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"kusnandartoni/starter/pkg/mail"
)

// EmailData :
type EmailData struct {
	EmailType string `json:"email_type"`
	Data      string `json:"data"`
}

func generateEmailTemplate(h mail.Mail, email mail.Format, templateName string) {
	pathFile := "runtime/mail"
	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(pathFile, 0744)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%v.html", pathFile, templateName), []byte(res), 0644)
	if err != nil {
		log.Fatal(err)
	}

	res, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%v.txt", pathFile, templateName), []byte(res), 0644)
	if err != nil {
		panic(err)
	}

}
