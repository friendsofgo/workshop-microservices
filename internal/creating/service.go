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
}

// NewService creates a creating service with the necessary dependencies
func NewService(cR counters.CounterRepository) Service {
	return &service{repository: cR}
}

// CreateCounter creates a new counter into your storage system
func (s service) CreateCounter(ctx context.Context, name, belongsTo string) error {
	c := counters.NewCounter(name, belongsTo)
	if err := s.repository.Save(ctx, c); err != nil {
		return err
	}
	return nil
}
