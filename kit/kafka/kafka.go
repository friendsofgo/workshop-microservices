package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type Dialer struct {
	kafkaDialer *kafka.Dialer
	brokers     []string
}

func Dial(brokers []string) *Dialer {
	dialer := &kafka.Dialer{
		ClientID: "workshop-microservices",
		Timeout:  10 * time.Second,
	}

	return &Dialer{kafkaDialer: dialer, brokers: brokers}
}
