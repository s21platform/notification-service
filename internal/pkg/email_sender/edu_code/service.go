package edu_code

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"

	"github.com/s21platform/notification-service/internal/config"
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

// Data - структура данных для шаблона
type Data struct {
	Code string
}

// SendVerificationCode отправляет верификационный код на указанный email
func (s *Service) SendVerificationCode(email string, code string) error {
	log.Printf("Sending verification code to %s", email)
	// Заполняем данные для шаблона
	data := Data{
		Code: code,
	}

	// Парсим шаблон письма
	tmpl, err := template.ParseFiles("templates/verification_edu.html")
	if err != nil {
		log.Printf("failed to parse template: %v", err)
		return err
	}

	// Заполняем шаблон данными
	var content bytes.Buffer
	if err := tmpl.Execute(&content, data); err != nil {
		log.Printf("failed to execute template: %v", err)
		return err
	}

	// Создаем и отправляем письмо
	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Код подтверждения Space-21")
	m.SetBody("text/html", content.String())

	log.Printf("try send: from %s, to %s", s.user, email)
	d := gomail.NewDialer(s.server, s.port, s.user, s.password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	//d.SSL = true
	if err := d.DialAndSend(m); err != nil {
		log.Printf("failed to send email: %v", err)
		return err
	}

	log.Printf("verification code for edu sent successfully to %s", email)
	return nil
}
