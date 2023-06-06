package sendgrid

import (
	"bytes"
	"fmt"
	"html/template"
	registerDto "iswift-go-project/internal/register/dto"
	"os"
	"path/filepath"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail interface {
	SendVerifivationEmail(toEmail string, dto registerDto.CreateEmailVerification)
}

type MailImpl struct {
}

func (mi *MailImpl) sendMail(toEmail, result, subject string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	to := mail.NewEmail(toEmail, toEmail)

	messages := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("MAIL_KEY"))

	resp, err := client.Send(messages)

	if err != nil {
		fmt.Println(err)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Println("email berhasil dikirim ke %s", toEmail)
	}
}

// SendVerifivationEmail implements Mail
func (mi *MailImpl) SendVerifivationEmail(toEmail string, dto registerDto.CreateEmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := ParseTemplate(templateFile, dto)

	if err != nil {
		fmt.Println(err)
	}

	mi.sendMail(toEmail, result, dto.SUBJECT)
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewMail() Mail {
	return &MailImpl{}
}
