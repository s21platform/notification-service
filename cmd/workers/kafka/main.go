package main

import (
	"context"
	"log"

	kafkalib "github.com/s21platform/kafka-lib"
	"github.com/s21platform/metrics-lib/pkg"

	"notification-service/internal/client/user"
	"notification-service/internal/config"
	"notification-service/internal/databus/invite_on_platform"
	"notification-service/internal/service/email_sender/invite_mail"
)

func main() {
	cfg := config.MustLoad()
	metrics, err := pkg.NewMetrics(cfg.Metrics.Host, cfg.Metrics.Port, "notification", cfg.Platform.Env)
	if err != nil {
		log.Fatalf("faild to connect graphite: %v", err)
	}

	ctx := context.WithValue(context.Background(), config.KeyMetrics, metrics)

	userClient := user.New(cfg)

	newFriendsConsumer, err := kafkalib.NewConsumer(cfg.Kafka.Server, cfg.Kafka.NotificationNewFriendTopic, metrics)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	inviteMail := invite_mail.New(cfg)

	inviteMailHandler := invite_on_platform.New(inviteMail, userClient)

	newFriendsConsumer.RegisterHandler(ctx, inviteMailHandler.Handle)

	log.Println("starting server")
	<-ctx.Done()
}
