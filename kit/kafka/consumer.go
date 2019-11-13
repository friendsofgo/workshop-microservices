package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// Handler defines a handle to execute an action after cossuming a message
type Handler func(ctx context.Context, msg []byte) error

// Consumer defines the properties of a publisher
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer creates a Consumer for kafka with the necessary dependencies
func NewConsumer(dialer *Dialer, topic, groupID string) *Consumer {
	c := kafka.ReaderConfig{
		Brokers:         dialer.brokers,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		GroupID:         groupID,
		StartOffset:     kafka.LastOffset,
	}

	return &Consumer{reader: kafka.NewReader(c)}
}

func (c *Consumer) Read(ctx context.Context, handle Handler) error {
	defer c.reader.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := c.handleMessage(ctx, handle); err != nil {
				return err
			}
		}
	}
}

func (c *Consumer) handleMessage(ctx context.Context, handle Handler) error {
	m, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return fmt.Errorf("error while reading a message: %v", err)
	}
	if err := handle(ctx, m.Value); err != nil {
		return err
	}

	return nil
}
