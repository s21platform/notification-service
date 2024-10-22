package invite_on_platform

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"log"
	"strings"

	"notification-service/internal/config"

	"github.com/s21platform/metrics-lib/pkg"
	"github.com/s21platform/notification-proto/notification-proto/new_user_invite"
)

type Email struct {
	User      string
	NewMember string
}

type Handler struct {
	imC InviteMailClient
	uC  UserClient
}

func New(imC InviteMailClient, uC UserClient) *Handler {
	return &Handler{imC: imC, uC: uC}
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

	tmpl, err := template.ParseFiles("templates/invite_on_platform.html")
	if err != nil {
		log.Fatal(err)
	}

	login, err := h.uC.GetLoginByUuid(ctx, msg.Uuid)
	if err != nil {
		m.Increment("new_friend.error")
		log.Printf("failed to get login: %v", err)
		return
	}
	data := Email{
		NewMember: strings.Split(msg.Email, "@")[0],
		User:      login,
	}

	var tmp bytes.Buffer
	err = tmpl.Execute(&tmp, data)
	if err != nil {
		log.Fatal(err)
	}

	err = h.imC.SendEmail("Приглашение на Space 21", msg.Email, tmp.String())
	if err != nil {
		m.Increment("new_friend.error")
		log.Printf("failed to send email: %v", err)
		return
	}
	m.Increment("new_friend.success")
	log.Println("send email success")
}
