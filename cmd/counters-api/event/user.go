package event

import (
	"context"
	"log"

	"github.com/friendsofgo/workshop-microservices/internal/creating"
	"github.com/friendsofgo/workshop-microservices/kit/domain"
)

const userCreatedEventType = "USER_CREATED"

type User struct {
	creatingService creating.Service
}

func NewUserHandler(creatingService creating.Service) *User {
	return &User{creatingService: creatingService}
}

type CreatedUserEventPayload struct {
	UserID   string `mapstructure:"user_id"`
	UserName string `mapstructure:"user_name"`
}

func (u *User) Handle(ctx context.Context, message []byte) error {
	var payload CreatedUserEventPayload
	m, err := domain.EventDecode(message, &payload)

	if err != nil {
		return err
	}

	switch m.EventType {
	case userCreatedEventType:
		log.Printf("%s message(%s) consumed", userCreatedEventType, m.ID)
		return u.creatingService.CreateCounter(ctx, "My first counter", payload.UserID)
	default:
		return nil
	}
}
