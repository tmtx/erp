package app

import (
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

// TODO: redo
func init() {
	// register event types
	// bus.RegisterMessageParamType(GuestCreated, &CreateGuestParams{})
	// bus.RegisterMessageParamType(UserCreated, &CreateUserParams{})
	// bus.RegisterMessageParamType(UserLoggedIn, &LoginUserParams{})
}

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
	Create(params CreateGuestParams) (*validator.Messages, error)
}

type UserService interface {
	Login(params LoginUserParams) (*validator.Messages, error)
}

// Command params

type CreateGuestParams struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type CreateUserParams struct {
	Email          string `bson:"email"`
	HashedPassword string `bson:"hashed_password"`
}

type LoginUserParams struct {
	Email          string
	HashedPassword string
}

// Aggregate

type Aggregate interface {
	Validate() (bool, *validator.Messages)
	ApplyEvent(e *event.Event)
}
