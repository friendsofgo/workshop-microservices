package counters

import (
	"context"

	"github.com/friendsofgo/workshop-microservices/kit/domain"
)

type publisher interface {
	Publish(ctx context.Context, payload interface{}) error
}

type Publisher struct {
	publisher publisher
}

func NewPublisher(p publisher) *Publisher {
	return &Publisher{publisher: p}
}

func (e *Publisher) Publish(ctx context.Context, events []domain.Event) error {
	for _, event := range events {
		if err := e.publisher.Publish(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
