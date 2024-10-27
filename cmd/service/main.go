package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"

	"notification-service/internal/config"
	"notification-service/internal/infra"
	"notification-service/internal/repository/postgres"
	"notification-service/internal/rpc"
)

func main() {
	cfg := config.MustLoad()
	dbRepo := postgres.New(cfg)
	defer dbRepo.Close()

	server := rpc.New(dbRepo)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			infra.AuthInterceptor,
		),
	)

	notificationproto.RegisterNotificationServiceServer(s, server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Service.Port))
	if err != nil {
		log.Printf("Cannot listen port: %s; Error: %s", cfg.Service.Port, err)
	}

	fmt.Println("Service started")
	if err = s.Serve(lis); err != nil {
		log.Printf("Cannot start service: %s; Error: %s", cfg.Service.Port, err)
	}
}
