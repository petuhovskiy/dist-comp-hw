package sendmail

import (
	"email/config"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	from   string
	dialer *gomail.Dialer
}

func NewSender(cfg config.Email) *Sender {
	return &Sender{
		from:   cfg.From,
		dialer: gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password),
	}
}

func (s *Sender) SendEmail(to string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "sample notification")
	m.SetBody("text/html", body)

	return s.dialer.DialAndSend(m)
}
