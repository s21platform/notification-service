package main

import (
	"fmt"
	"notification-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Service)
}
