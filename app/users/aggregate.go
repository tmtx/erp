package users

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type User struct {
	Id             *event.UUID
	Email          string
	HashedPassword string
}

func (u *User) ApplyEvent(e *event.Event) {
	switch e.Key {
	case app.UserCreated:
		params := e.Params.(app.CreateGuestParams)

		u.Email = params.Email
		u.Id = e.EntityId
	}
}

func (u *User) Validate() (bool, *validator.Messages) {
	return false, nil
}
