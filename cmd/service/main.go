package main

import (
	"fmt"
	"log"
	"notification-service/internal/config"
	"notification-service/internal/repository/postgres"
)

func main() {
	cfg := config.MustLoad()
	dbRepo, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("db.New: %w", err))
	}
	defer dbRepo.Close()
}
