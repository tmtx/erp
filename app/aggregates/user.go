package aggregates

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type User struct {
	Base
	Email          string `bson:"email"`
	HashedPassword string `bson:"hashed_password"`
}

func (ag *User) GetTargetEvents() []bus.MessageKey {
	return []bus.MessageKey{
		app.UserCreated,
		// app.UserInfoUpdated,
	}
}

func (ag *User) HydrateFromParams(params bus.MessageParams) {
	if email, ok := params["email"].(string); ok {
		ag.Email = email
	}
	if pw, ok := params["hashed_password"].(string); ok {
		ag.HashedPassword = pw
	}
}

func (ag *User) ApplyEvent(e event.Event) {
	switch e.Key {
	case app.UserCreated:
		fallthrough
	case app.UserInfoUpdated:
		ag.HydrateFromParams(e.Params)
		if e.EntityId != nil {
			ag.Id = e.EntityId
		}
	}
}

func (ag *User) Validate() (bool, *validator.Messages) {
	return false, nil
}

func (ag *User) CanBeRestored() bool {
	if ag.Id == nil && ag.Email == "" {
		return false
	}

	return true
}

func (ag *User) Restore() error {
	if !ag.CanBeRestored() {
		return app.TypedError{Type: app.UnableToRestoreAggregate}
	}

	if ag.Id != nil {
		restoredAggregate, err := RestoreFromId(ag)
		if err != nil {
			return err
		}
		ag = restoredAggregate.(*User)
	} else if ag.Email != "" {
		restoredAggregate, err := RestoreFromEmail(ag, ag.Email)
		if err != nil {
			return err
		}
		ag = restoredAggregate.(*User)

	}

	return nil
}
