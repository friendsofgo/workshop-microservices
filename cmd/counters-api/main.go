package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/event"
	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/server/http"
	counters "github.com/friendsofgo/workshop-microservices/internal"
	"github.com/friendsofgo/workshop-microservices/internal/creating"
	"github.com/friendsofgo/workshop-microservices/internal/fetching"
	"github.com/friendsofgo/workshop-microservices/internal/storage/mongo"
	"github.com/friendsofgo/workshop-microservices/kit/kafka"

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

		brokersStr   = os.Getenv("WORKSHOP_KAFKA_BROKERS")
		brokers      = strings.Split(brokersStr, ",")
		userTopic    = os.Getenv("WORKSHOP_KAFKA_USER_TOPIC")
		userGroup    = os.Getenv("WORKSHOP_KAFKA_USER_GROUP")
		counterTopic = os.Getenv("WORKSHOP_KAFKA_COUNTER_TOPIC")
	)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, err := mongo.Dial(rootCtx, mongoHost, mongoPort)
	if err != nil {
		log.Fatalln("error trying to connect with mongo", err)
	}

	var (
		dialer = kafka.Dial(brokers)
		kafkaCounterPublisher = kafka.NewPublisher(dialer, counterTopic)
		// maybe you need this...
		_                     = counters.NewPublisher(kafkaCounterPublisher)

		counterRepository = mongo.NewCounterRepository(mongoClient.Database(mongoDB))
		creatingService   = creating.NewService(counterRepository)
		fetchingService   = fetching.NewService(counterRepository)

		userEventHandler = event.NewUserHandler(creatingService)
		userConsumer     = kafka.NewConsumer(dialer, userTopic, userGroup)
	)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := userConsumer.Read(rootCtx, userEventHandler.Handle); err != nil {
			cancel()
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()
		srv := http.NewServer(context.Background(), _defaultHost, _defaultPort, creatingService, fetchingService, logger)
		if err := srv.Serve(); err != nil {
			cancel()
			log.Fatalln(err)
		}
	}()

	wg.Wait()
}
