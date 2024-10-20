package user_invite

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"log"

	"notification-service/internal/config"

	"github.com/s21platform/metrics-lib/pkg"
	"github.com/s21platform/notification-proto/notification-proto/new_user_invite"
)

type Email struct {
	From string
}

type Handler struct {
	imC InviteMailClient
}

func New(imC InviteMailClient) *Handler {
	return &Handler{imC: imC}
}

func convertMessage(bMessage []byte, target interface{}) error {
	err := json.Unmarshal(bMessage, target)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, in []byte) {
	m := pkg.FromContext(ctx, config.KeyMetrics)

	var msg new_user_invite.NewUserInvite
	err := convertMessage(in, &msg)
	if err != nil {
		m.Increment("new_friend.error")
		log.Printf("failed to convert message: %v", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/welcome.html")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("prepare tmpl")
	data := Email{
		From: "garroshm",
	}

	var tmp bytes.Buffer
	err = tmpl.Execute(&tmp, data)
	if err != nil {
		log.Fatal(err)
	}

	err = h.imC.SendEmail("Приглашение на Space 21", msg.Email, msg.String())
	if err != nil {
		m.Increment("new_friend.error")
		log.Printf("failed to send email: %v", err)
		return
	}
	m.Increment("new_friend.success")
	log.Println("send email success")
}
