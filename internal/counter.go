package counters

import (
	"context"
	"time"

	"github.com/friendsofgo/workshop-microservices/kit/domain"
	"github.com/friendsofgo/workshop-microservices/kit/ulid"
)

// Counter counter entity
type Counter struct {
	ID        string     `bson:"_id" json:"id"`
	Name      string     `bson:"name" json:"name"`
	Value     uint       `bson:"value" json:"value"`
	BelongsTo string     `bson:"belongs_to" json:"belongs_to"`
	CreatedAt time.Time  `bson:"created_at" json:"-"`
	UpdatedAt *time.Time `bson:"updated_at" json:"-"`

	events []domain.Event `bson:"-" json:"-"`
}

// NewCounter initialize a counter entity
func NewCounter(name, belongsTo string) *Counter {
	c := &Counter{
		ID:        ulid.New(),
		Name:      name,
		BelongsTo: belongsTo,
	}
	return c
}

// Record store a new event in the structure
func (c *Counter) Record(evt domain.Event) {
	// code here
}

func (c Counter) Events() []domain.Event {
	return c.events
}

// CounterRepository declare the necessary interface to our repository
type CounterRepository interface {
	FetchAllByUser(ctx context.Context, belongsTo string) ([]Counter, error)
	Save(ctx context.Context, counter *Counter) error
}
