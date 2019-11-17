package creating

import (
	"context"

	counters "github.com/friendsofgo/workshop-microservices/internal"
)

// Service provides counter creating operations.
type Service interface {
	CreateCounter(ctx context.Context, name, belongsTo string) error
}

type service struct {
	repository counters.CounterRepository
	publisher  *counters.Publisher
}

// NewService creates a creating service with the necessary dependencies
func NewService(cR counters.CounterRepository, p *counters.Publisher) Service {
	return &service{repository: cR, publisher: p}
}

// CreateCounter creates a new counter into your storage system
func (s service) CreateCounter(ctx context.Context, name, belongsTo string) error {
	c := counters.NewCounter(name, belongsTo)
	if err := s.repository.Save(ctx, c); err != nil {
		return err
	}

	if err := s.publisher.Publish(ctx, c.Events()); err != nil {
		return err
	}
	return nil
}
