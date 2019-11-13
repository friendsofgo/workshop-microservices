package event

import (
	"context"
	"encoding/json"

	"github.com/friendsofgo/workshop-microservices/internal/creating"
)

const userCreatedEventType = "USER_CREATED"

type User struct {
	creatingService creating.Service
}

func NewUserHandler(creatingService creating.Service) *User {
	return &User{creatingService: creatingService}
}

type UserMessage struct {
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	Data      struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
	} `json:"data"`
}

func (u *User) Handle(ctx context.Context, message []byte) error {
	m, err := u.decodeMessage(message)
	if err != nil {
		return err
	}

	switch m.EventType {
	case userCreatedEventType:
		return u.creatingService.CreateCounter(ctx, "My first counter", m.Data.UserID)
	default:
		return nil
	}
}

func (u *User) decodeMessage(message []byte) (UserMessage, error) {
	var decoded UserMessage
	err := json.Unmarshal(message, &decoded)
	return decoded, err
}
