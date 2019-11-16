package fetching

import (
	"context"

	counters "github.com/friendsofgo/workshop-microservices/internal"
)

// Service provides counter fetching operations
type Service interface {
	FetchAllCountersByUser(ctx context.Context, belongsTo string) ([]counters.Counter, error)
}

type service struct {
	repository counters.CounterRepository
}

// NewService creates a fetching service with the necessary dependencies
func NewService(cR counters.CounterRepository) Service {
	return &service{repository: cR}
}

// CreateCounter creates a new counter into your storage system
func (s service) FetchAllCountersByUser(ctx context.Context, belongsTo string) ([]counters.Counter, error) {
	return s.repository.FetchAllByUser(ctx, belongsTo)
}
