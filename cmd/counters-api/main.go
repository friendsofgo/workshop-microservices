package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/server/http"
	"github.com/friendsofgo/workshop-microservices/internal/creating"
	"github.com/friendsofgo/workshop-microservices/internal/storage/mongo"

	_ "github.com/joho/godotenv/autoload"
)

const (
	_defaultHost      = "localhost"
	_defaultPort uint = 3000
)

func main() {

	var (
		mongoHost    = os.Getenv("WORKSHOP_MONGO_HOST")
		mongoPort, _ = strconv.ParseUint(os.Getenv("WORKSHOP_MONGO_PORT"), 10, 32)
		mongoDB      = os.Getenv("WORKSHOP_MONGO_DB")
	)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, err := mongo.Dial(rootCtx, mongoHost, mongoPort)
	if err != nil {
		log.Fatalln("error trying to connect with mongo", err)
	}

	var (
		counterRepository = mongo.NewCounterRepository(mongoClient.Database(mongoDB))
		creatingService   = creating.NewService(counterRepository)
	)

	srv := http.NewServer(context.Background(), _defaultHost, _defaultPort, creatingService, logger)
	if err := srv.Serve(); err != nil {
		cancel()
		log.Fatalln(err)
	}
}
