package event

import (
	"context"
	"encoding/json"
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
	m, err := u.decodeMessage(message)
	if err != nil {
		return err
	}

	switch m.EventType {
	case userCreatedEventType:
		log.Printf("%s message(%s) consumed", userCreatedEventType, m.ID)
		var payload CreatedUserEventPayload
		err := m.DecodePayload(&payload)
		if err != nil {
			return err
		}

		return u.creatingService.CreateCounter(ctx, "My first counter", payload.UserID)
	default:
		return nil
	}
}

func (u *User) decodeMessage(message []byte) (domain.Event, error) {
	var decoded domain.Event
	err := json.Unmarshal(message, &decoded)
	return decoded, err
}
