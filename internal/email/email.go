package email

import (
	"bytes"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

type SmtpConfig struct {
	SmtpHost string `env:"SMTP_HOST,required"`
	SmtpPort int    `env:"SMTP_PORT,required"`
	Username string `env:"SMTP_USER,required"`
	Password string `env:"SMTP_PASS,required"`
}

// sends email message via SMTP
func (smtpCfg *SmtpConfig) Send(sender, recipient, message string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", recipient)
	msg.SetHeader("Subject", "🎅🏻 Ho Ho Ho, your Secret Santa is here!")
	msg.SetBody("text/html", message)

	fmt.Printf("Sending email to: %s\n", recipient)

	d := gomail.NewDialer(smtpCfg.SmtpHost, smtpCfg.SmtpPort, smtpCfg.Username, smtpCfg.Password)
	return d.DialAndSend(msg)
}

type EmailData struct {
	Name        string
	SecretSanta string
	Message     string
}

// parses the given HTML template and fills it with EmailData
func ParseTemplate(templatePath string, data EmailData) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
