package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

func (c counterRepository) FetchAllByUser(ctx context.Context, belongsTo string) ([]counters.Counter, error) {
	findOpts := options.Find()
	findOpts.SetSort(bson.M{"created_at": -1})

	cursor, err := c.db.Collection(c.collection).Find(ctx, bson.M{"belongs_to": belongsTo}, findOpts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var list []counters.Counter
	for cursor.Next(ctx) {
		var counter counters.Counter
		if err := cursor.Decode(&counter); err != nil {
			continue
		}

		list = append(list, counter)
	}

	return list, nil
}
