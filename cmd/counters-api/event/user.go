package event

import (
	"context"

	"github.com/friendsofgo/workshop-microservices/internal/creating"
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
	// code here
	return nil
}
