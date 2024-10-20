package main

import (
	"context"
	"fmt"
	"log"

	"notification-service/internal/config"
	"notification-service/internal/databus/user_invite"
	"notification-service/internal/service/email_sender/invite_mail"

	kafkalib "github.com/s21platform/kafka-lib"
	"github.com/s21platform/metrics-lib/pkg"
)

type Email struct {
	Name string
	Code string
}

func main() {
	cfg := config.MustLoad()
	//dbRepo, err := postgres.New(cfg)
	//if err != nil {
	//	log.Fatal(fmt.Errorf("db.New: %w", err))
	//}
	//defer dbRepo.Close()
	fmt.Println(cfg.EmailVerification.Server, cfg.EmailVerification.Port, cfg.EmailVerification.User, cfg.EmailVerification.Password)
	metrics, err := pkg.NewMetrics(cfg.Metrics.Host, cfg.Metrics.Port, "notification", cfg.Platform.Env)
	if err != nil {
		log.Fatalf("faild to connect graphite: %v", err)
	}

	ctx := context.WithValue(context.Background(), config.KeyMetrics, metrics)

	newFriendsConsumer, err := kafkalib.NewConsumer(cfg.Kafka.Server, cfg.Kafka.NotificationNewFriendTopic, metrics)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	inviteMail := invite_mail.New(cfg)

	inviteMailHandler := user_invite.New(inviteMail)

	newFriendsConsumer.RegisterHandler(ctx, inviteMailHandler.Handle)

	<-ctx.Done()
}
