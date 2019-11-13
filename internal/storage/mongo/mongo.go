package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dial generate a connection with mongodb
func Dial(ctx context.Context, host string, port uint64) (*mongo.Client, error) {
	mgoAddr := fmt.Sprintf("mongodb://%s:%d", host, port)
	opts := options.Client().ApplyURI(mgoAddr)

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	//Check the connections
	err = c.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}
