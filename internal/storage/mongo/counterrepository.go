package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	counters "github.com/friendsofgo/workshop-microservices/internal"
)

type counterRepository struct {
	db         *mongo.Database
	collection string
}

func NewCounterRepository(db *mongo.Database) counters.CounterRepository {
	return &counterRepository{
		db:         db,
		collection: "counters",
	}
}

func (c counterRepository) Save(ctx context.Context, counter *counters.Counter) error {
	counter.CreatedAt = time.Now()
	if _, err := c.db.Collection(c.collection).InsertOne(ctx, counter); err != nil {
		return err
	}

	return nil
}
