package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/s21platform/notification-service/pkg/notification"

	"github.com/s21platform/notification-service/internal/config"
	"github.com/s21platform/notification-service/internal/infra"
	"github.com/s21platform/notification-service/internal/repository/postgres"
	"github.com/s21platform/notification-service/internal/service"
)

func main() {
	cfg := config.MustLoad()
	dbRepo := postgres.New(cfg)
	defer dbRepo.Close()

	server := service.New(dbRepo)

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
