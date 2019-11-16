package counters

import (
	"context"
	"time"

	"github.com/friendsofgo/workshop-microservices/kit/ulid"
)

// Counter counter entity
type Counter struct {
	ID        string     `bson:"_id" json:"id"`
	Name      string     `bson:"name" json:"name"`
	Value     uint       `bson:"value" json:"value"`
	BelongsTo string     `bson:"belongs_to" json:"belongs_to"`
	CreatedAt time.Time `bson:"created_at" json:"-"`
	UpdatedAt *time.Time `bson:"updated_at" json:"-"`
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
	FetchAllByUser(ctx context.Context, belongsTo string) ([]Counter, error)
	Save(ctx context.Context, counter *Counter) error
}
