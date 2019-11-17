package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"

	"github.com/friendsofgo/workshop-microservices/kit/ulid"
)

// Publisher defines the properties of a publisher
type Publisher struct {
	writer *kafka.Writer
}

// NewPublisher creates a Publisher for kafka with the necessary dependencies
func NewPublisher(dialer *Dialer, topic string) *Publisher {
	c := kafka.WriterConfig{
		Brokers:          dialer.brokers,
		Topic:            topic,
		Dialer:           dialer.kafkaDialer,
		Balancer:         &kafka.LeastBytes{},
		ReadTimeout:      10 * time.Second,
		WriteTimeout:     10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return &Publisher{kafka.NewWriter(c)}
}

// Publish send a message into a kafka topic
func (p *Publisher) Publish(ctx context.Context, payload interface{}) error {
	message, err := p.encodeMessage(payload)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, message)
}

func (p *Publisher) encodeMessage(payload interface{}) (kafka.Message, error) {
	m, err := json.Marshal(payload)
	if err != nil {
		return kafka.Message{}, err
	}

	key := ulid.New()
	return kafka.Message{
		Key:   []byte(key),
		Value: m,
	}, nil
}
