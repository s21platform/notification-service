package main

import (
	"notification-service/internal/config"
)

func main() {
	_ = config.MustLoad()
}
