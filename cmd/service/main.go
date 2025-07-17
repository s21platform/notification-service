package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/s21platform/notification-service/internal/config"
	"github.com/s21platform/notification-service/internal/infra"
	"github.com/s21platform/notification-service/internal/pkg/email_sender"
	"github.com/s21platform/notification-service/internal/pkg/email_sender/edu_code"
	"github.com/s21platform/notification-service/internal/pkg/email_sender/verification_code"
	"github.com/s21platform/notification-service/internal/repository/postgres"
	"github.com/s21platform/notification-service/internal/service"
	"github.com/s21platform/notification-service/pkg/notification"
)

func main() {
	cfg := config.MustLoad()
	db := postgres.New(cfg)
	defer db.Close()

	// Инициализируем email сервис
	emailSender := email_sender.New(cfg)

	// Инициализируем сервис верификационных кодов
	verificationCodeSender := verification_code.New(cfg)
	vecS := edu_code.New(cfg)

	server := service.New(db, emailSender, verificationCodeSender, vecS)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			infra.AuthInterceptor,
		),
	)

	notification.RegisterNotificationServiceServer(s, server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Service.Port))
	if err != nil {
		log.Printf("Cannot listen port: %s; Error: %s", cfg.Service.Port, err)
	}

	fmt.Printf("Service started on port: %s\n", cfg.Service.Port)
	if err = s.Serve(lis); err != nil {
		log.Printf("Cannot start service: %s; Error: %s", cfg.Service.Port, err)
	}
}
