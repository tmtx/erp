package aggregates

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

type User struct {
	Base
	Email          string `bson:"email"`
	HashedPassword string `bson:"hashed_password"`
}

func (ag *User) GetTargetEvents() []bus.MessageKey {
	return []bus.MessageKey{
		app.UserCreated,
		app.UserInfoUpdated,
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
	case app.UserInfoUpdated:
		fallthrough
	case app.UserCreated:
		ag.HydrateFromParams(e.Params)
		if e.EntityId != nil {
			ag.Id = e.EntityId
		}
	}
	ag.AppliedEventCount += 1
}

func (ag *User) Validate() (bool, *validator.Messages) {
	if ag.HashedPassword == "" {
		return false, &validator.Messages{
			"hashed_password": []validator.Message{
				"empty password",
			},
		}
	}
	if ag.AppliedEventCount != uint(len(ag.GetEvents())) {
		return false, &validator.Messages{
			"events": []validator.Message{
				"not all events applied",
			},
		}
	}
	return true, nil
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
