package invite_mail

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"

	"notification-service/internal/config"
)

type Service struct {
	server   string
	port     int
	user     string
	password string
	env      string
}

func New(cfg *config.Config) *Service {
	return &Service{
		server:   cfg.EmailVerification.Server,
		port:     cfg.EmailVerification.Port,
		user:     cfg.EmailVerification.User,
		password: cfg.EmailVerification.Password,
		env:      cfg.Platform.Env,
	}
}

func (s *Service) SendEmail(subject string, to string, content string) error {
	if s.env != "prod" {
		whiteList := map[string]struct{}{
			"garroshm@student.21-school.ru": {},
		}
		if _, ok := whiteList[to]; !ok {
			return fmt.Errorf("invalid email address for this environment")
		}
	}

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
		log.Printf("failed to send email: %v", err)
		return err
	}
	return nil
}
