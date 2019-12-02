package app

import (
	"github.com/google/uuid"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

// Commands
const (
	CreateGuest bus.MessageKey = "create_guest"
	CreateUser  bus.MessageKey = "create_user"
	LoginUser   bus.MessageKey = "login_user"
)

// Events
const (
	GuestCreated bus.MessageKey = "guest_created"
	UserLoggedIn bus.MessageKey = "user_logged_in"
	UserCreated  bus.MessageKey = "user_created"
)

type UUID struct {
	uuid.UUID
}

type Server interface {
	Start(conn string)
}

type Routable interface {
	NewRouter()
}

type CommandSubscriber interface {
	RegisterCommandCallbacks()
}

// Services

type GuestService interface {
	Create(params CreateGuestParams) (error, *validator.Messages)
}

type UserService interface {
	Login(params LoginUserParams) (error, *validator.Messages)
}

// Command params

type CreateGuestParams struct {
	Id    UUID   `bson:"id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type CreateUserParams struct {
	Id             UUID
	Email          string
	HashedPassword string
}

type LoginUserParams struct {
	Email          string
	HashedPassword string
}

// Aggregate

type Aggregate interface {
	Validate() (error, *validator.Messages)
	ApplyEvent(e event.Event)
}

type AggregateRootRepository interface {
	RestoreAggregateRootById(id UUID) (error, *Aggregate)
}
