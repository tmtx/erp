package app

import (
	"time"

	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

// Commands
// TODO: rename CommandCreateGuest etc.
const (
	CreateGuest       bus.MessageKey = "create_guest"
	CreateUser        bus.MessageKey = "create_user"
	LoginUser         bus.MessageKey = "login_user"
	UpdateUserInfo    bus.MessageKey = "update_user_info"
	CreateReservation bus.MessageKey = "create_reservation"
)

// Events
// TODO: rename EventGuestCreted etc.
const (
	GuestCreated       bus.MessageKey = "guest_created"
	UserLoggedIn       bus.MessageKey = "user_logged_in"
	UserCreated        bus.MessageKey = "user_created"
	UserInfoUpdated    bus.MessageKey = "user_info_updated"
	ReservationCreated bus.MessageKey = "reservation_created"
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
	Create(params CreateGuestParams) (validator.Messages, error)
	RestoreAggregateRootByEmail(email string) (Aggregate, error)
}

type UserService interface {
	Login(params LoginUserParams) (validator.Messages, error)
	UpdateUserInfo(params UpdateUserInfoParams) (validator.Messages, error)
	Session(sessValues interface{}) server.Session
}

type SpacesService interface {
	GetAllAvailable() ([]Aggregate, error)
}

type ReservationsService interface {
	Create(params CreateReservationParams) (validator.Messages, error)
	GetAll() ([]Aggregate, error)
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

type CreateReservationParams struct {
	GuestName  string
	GuestEmail string
	StartDate  time.Time
	EndDate    time.Time
	SpaceId    uint
}

// Aggregate

type Aggregate interface {
	Validate() (bool, *validator.Messages)
	ApplyEvent(e event.Event)
	CanBeRestored() bool
	Restore() error
	GetId() *event.UUID
	GetRepository() event.Repository
	GetTargetEvents() []bus.MessageKey
}
