package aggregates

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type Guest struct {
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ag *Guest) GetTargetEvents() []bus.MessageKey {
	return []bus.MessageKey{
		app.GuestCreated,
	}
}

func (ag *Guest) ApplyEvent(e event.Event) {
	switch e.Key {
	case app.GuestCreated:
		params := e.Params

		if params["email"] != nil {
			ag.Email = params["email"].(string)
		}
		if params["name"] != nil {
			ag.Name = params["name"].(string)
		}
		ag.Id = e.EntityId
	case app.ReservationCreated:
		params := e.Params

		if params["guest_email"] != nil {
			ag.Email = params["guest_email"].(string)
		}
		if params["guest_name"] != nil {
			ag.Name = params["guest_name"].(string)
		}
	}
}

func (ag *Guest) Validate() (bool, *validator.Messages) {
	// TODO: implement
	return true, nil
}

func (ag *Guest) CanBeRestored() bool {
	if ag.Id == nil && ag.Email == "" {
		return false
	}
	return true
}

func (ag *Guest) Restore() error {
	if !ag.CanBeRestored() {
		return app.TypedError{Type: app.UnableToRestoreAggregate}
	}

	if ag.Id != nil {
		restoredAggregate, err := RestoreFromId(ag)
		if err != nil {
			return err
		}
		ag = restoredAggregate.(*Guest)

	} else if ag.Email != "" {
		restoredAggregate, err := RestoreFromEmail(ag, ag.Email)
		if err != nil {
			return err
		}
		ag = restoredAggregate.(*Guest)
	}

	return nil
}
