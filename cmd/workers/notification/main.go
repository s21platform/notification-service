package main

import (
	"notification-service/internal/repository/kafka"
)

func main() {
	kafka.Consume()
}
