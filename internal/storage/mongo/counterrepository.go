package mongo

import (
	"context"
	counters "github.com/friendsofgo/workshop-microservices/internal"
	"go.mongodb.org/mongo-driver/mongo"
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
	if _, err := c.db.Collection(c.collection).InsertOne(ctx, counter); err != nil {
		return err
	}

	return nil
}
