package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

const (
	BROKER = "localhost:9092"
	TOPIC  = "test"
)

func Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     strings.Split(BROKER, "."),
		GroupID:     strconv.Itoa(rand.Intn(500)),
		Topic:       TOPIC,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
	})
	defer r.Close()
	var n Notification
	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			log.Fatal(err)

		}
		json.Unmarshal(m.Value, &n)
		fmt.Printf("получено сообщение: %s\n", n)

	}
}
