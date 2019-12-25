package app

import (
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

// Commands
const (
	CreateGuest    bus.MessageKey = "create_guest"
	CreateUser     bus.MessageKey = "create_user"
	LoginUser      bus.MessageKey = "login_user"
	UpdateUserInfo bus.MessageKey = "update_user_info"
)

// Events
const (
	GuestCreated    bus.MessageKey = "guest_created"
	UserLoggedIn    bus.MessageKey = "user_logged_in"
	UserCreated     bus.MessageKey = "user_created"
	UserInfoUpdated bus.MessageKey = "user_info_updated"
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
	Login(params LoginUserParams) (validator.Messages, error)
	UpdateUserInfo(params UpdateUserInfoParams) (validator.Messages, error)
	RestoreAggregateRootByEmail(email string) (Aggregate, error)
	Session(sessValues interface{}) server.Session
}

type SpacesService interface {
	GetAllAvailable() ([]Aggregate, error)
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
	Email    string
	Password string
}

type UpdateUserInfoParams struct {
	UserId *event.UUID
	Email  string
}

// Aggregate

type Aggregate interface {
	Validate() (bool, *validator.Messages)
	ApplyEvent(e event.Event)
}
