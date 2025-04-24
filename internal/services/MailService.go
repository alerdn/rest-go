package services

import (
	"github.com/alerdn/rest-go/config"
	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {config.SMTP.User},
		"To":      {to},
		"Subject": {subject},
	})
	m.SetBody("text/html", body)

	d := gomail.NewDialer(config.SMTP.Host, config.SMTP.Port, config.SMTP.User, config.SMTP.Pass)

	return d.DialAndSend(m)
}
