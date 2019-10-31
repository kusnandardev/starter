package mail

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig"
	"github.com/imdario/mergo"
	"github.com/jaytaylor/html2text"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// Mail :
type Mail struct {
	Theme    Theme
	MailType string
}

// Theme :
type Theme interface {
	Name() string
	HTMLTemplate() string
	PlainTextTemplate() string
}

// Forgot :
type Forgot struct {
	Name       string `json:"name"`
	ButtonLink string `json:"button_link"`
}

// Format :
type Format struct {
	Forgot Forgot
}

// Template :
type Template struct {
	Mail   Mail
	Format Format
}

var templateFuncs = template.FuncMap{
	"url": func(s string) template.URL {
		return template.URL(s)
	},
	"markDown": func(s string) template.HTML {
		unsafe := blackfriday.Run([]byte(s))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		return template.HTML(html)
	},
}

// GenerateHTML :
func (h *Mail) GenerateHTML(email Format) (string, error) {
	err := setDefaultMailValues(h)
	if err != nil {
		return "", err
	}
	return h.generateTemplates(email, h.Theme.HTMLTemplate())
}

// GeneratePlainText :
func (h *Mail) GeneratePlainText(email Format) (string, error) {
	err := setDefaultMailValues(h)
	if err != nil {
		return "", err
	}
	template, err := h.generateTemplates(email, h.Theme.PlainTextTemplate())
	if err != nil {
		return "", err
	}
	return html2text.FromString(template, html2text.Options{PrettyTables: true})
}

func (h *Mail) generateTemplates(email Format, tplt string) (string, error) {
	err := setDefaultFormatValues(&email)
	if err != nil {
		return "", err
	}
	t, err := template.New("himail").Funcs(sprig.FuncMap()).Funcs(templateFuncs).Parse(tplt)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	t.Execute(&b, Template{*h, email})
	return b.String(), nil
}

func setDefaultMailValues(h *Mail) error {
	defaultMail := Mail{
		Theme:    new(Default),
		MailType: "",
	}

	err := mergo.Merge(h, defaultMail)
	if err != nil {
		return err
	}
	return nil
}

func setDefaultFormatValues(e *Format) error {
	defaultFormat := Format{
		Forgot: Forgot{
			Name:       "",
			ButtonLink: "https://supergroup.helochat.id",
		},
		// Register: Register{
		// 	PICName: "",
		// 	ORGName: "",
		// 	ORGType: "",
		// },
		// ContactUs: ContactUs{
		// 	Name:    "",
		// 	Email:   "",
		// 	Subject: "",
		// 	Message: "",
		// },
		// Approve: Approve{
		// 	PICName:    "",
		// 	ORGName:    "",
		// 	ORGType:    "",
		// 	ButtonLink: "https://supergroup.helochat.id",
		// },
		// Reject: Reject{
		// 	PICName: "",
		// 	ORGName: "",
		// 	ORGType: "",
		// 	Remark:  "",
		// },
		// AdminInvite: AdminInvite{
		// 	Name:       "",
		// 	ButtonLink: "https://supergroup.helochat.id",
		// },
	}
	return mergo.Merge(e, defaultFormat)
}
