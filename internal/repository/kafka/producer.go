package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

type Notification struct {
	ID     string `json:"id"`
	UUID   string `json:"uuid"`
	Value  string `json:"value"`
	IsRead string `json:"isRead"`
}

func Produce() {
	w := &kafka.Writer{
		Addr:     kafka.TCP(strings.Split(BROKER, ".")...),
		Topic:    TOPIC,
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()

	n1 := Notification{"1", "1", "peer point", "true"}
	//n2 := Notification{"2", "2", "deadline", "false"}
	//n3 := Notification{"3", "5", "new task", "true"}
	buf, _ := json.Marshal(n1)

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(buf),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
