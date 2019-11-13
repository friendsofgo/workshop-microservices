package counters

import (
	"context"
	"time"

	"github.com/friendsofgo/workshop-microservices/kit/ulid"
)

// Counter counter entity
type Counter struct {
	ID        string     `bson:"_id"`
	Name      string     `bson:"name"`
	Value     uint       `bson:"value"`
	BelongsTo string     `bson:"belongs_to"`
	UpdatedAt *time.Time `bsong:"updated_at"`
}

// NewCounter initialize a counter entity
func NewCounter(name, belongsTo string) *Counter {
	return &Counter{
		ID:        ulid.New(),
		Name:      name,
		BelongsTo: belongsTo,
	}
}

// CounterRepository declare the necessary interface to our repository
type CounterRepository interface {
	Save(ctx context.Context, counter *Counter) error
}
