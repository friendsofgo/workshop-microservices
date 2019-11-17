package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/friendsofgo/workshop-microservices/kit/domain"
	"github.com/friendsofgo/workshop-microservices/kit/ulid"

	"github.com/friendsofgo/workshop-microservices/kit/kafka"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	createUserMessage := flag.Bool("create-user-msg", false, "create a message of type event user created")
	flag.Parse()
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if *createUserMessage {
		publishcreateUserMessage(rootCtx)
		return
	}

	flag.Usage()
}

func publishcreateUserMessage(ctx context.Context) {
	var (
		brokers = os.Getenv("WORKSHOP_KAFKA_BROKERS")
		topic   = os.Getenv("WORKSHOP_KAFKA_USER_TOPIC")
	)
	b := strings.Split(brokers, ",")

	dialer := kafka.Dial(b)
	publisher := kafka.NewPublisher(dialer, topic)

	userID := ulid.New()
	type data struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
	}
	userCreatedEvt := domain.Event{
		ID:          ulid.New(),
		EventType:   "USER_CREATED",
		AggregateID: userID,
		Payload: data{
			UserID:   userID,
			UserName: fmt.Sprintf("fogo_user_%s", userID),
		},
		OccurredOn: time.Now(),
	}
	if err := publisher.Publish(ctx, userCreatedEvt); err != nil {
		log.Fatalln("error trying to publish USER_CREATED message on kafka: ", err)
	}

	log.Printf("USER_CREATED event: %s published\n", userCreatedEvt.ID)
}
