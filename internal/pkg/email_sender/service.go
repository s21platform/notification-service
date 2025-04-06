package email_sender

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"

	"github.com/s21platform/notification-service/internal/config"
)

type Service struct {
	server   string
	port     int
	user     string
	password string
}

func New(cfg *config.Config) *Service {
	return &Service{
		server:   cfg.EmailVerification.Server,
		port:     cfg.EmailVerification.Port,
		user:     cfg.EmailVerification.User,
		password: cfg.EmailVerification.Password,
	}
}

func (s *Service) SendEmail(subject string, to string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(s.server, s.port, s.user, s.password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	//d.SSL = true
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
