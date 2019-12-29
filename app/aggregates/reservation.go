package aggregates

import (
	"time"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Base
	SpaceId   *uint `json:"spaceId"`
	Guest     `json:"guest"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

func (ag *Reservation) HydrateFromParams(params bus.MessageParams) {
	if email, ok := params["guest_email"].(string); ok {
		ag.Guest.Email = email
	}
	if name, ok := params["guest_name"].(string); ok {
		ag.Guest.Name = name
	}
	if spaceId, ok := params["space_id"].(int64); ok {
		spaceId := uint(spaceId)
		ag.SpaceId = &spaceId
	}
	if startDateTimestamp, ok := params["start_date"].(primitive.DateTime); ok {
		startDate := time.Unix(int64(int64(startDateTimestamp)/1000), 0)
		ag.StartDate = &startDate
	}
	if endDateTimestamp, ok := params["end_date"].(primitive.DateTime); ok {
		endDate := time.Unix(int64(int64(endDateTimestamp)/1000), 0)
		ag.EndDate = &endDate
	}
}

func (ag *Reservation) ApplyEvent(e event.Event) {
	switch e.Key {
	case app.ReservationCreated:
		params := e.Params
		ag.Guest.Base = Base{Repository: ag.Repository}
		ag.HydrateFromParams(params)
		if ag.Email != "" {
			ag.Guest.Restore()
		}
		if e.EntityId != nil {
			ag.Id = e.EntityId
		}
	}
}

func (ag *Reservation) Validate() (bool, *validator.Messages) {
	// TODO: implement
	return false, nil
}

func (ag *Reservation) CanBeRestored() bool {
	if ag.Guest.Email == "" || ag.SpaceId == nil {
		return false
	}
	return true
}

func (ag *Reservation) Restore() error {
	if !ag.CanBeRestored() {
		return app.TypedError{Type: app.UnableToRestoreAggregate}
	}

	restoredAg, err := RestoreWithFilter(ag, event.Filter{
		"key": event.Filter{"$in": ag.GetTargetEvents()},
		"params": event.Filter{
			"guest_email": ag.Guest.Email,
			"space_id":    ag.SpaceId,
		},
	})
	if err != nil {
		return err
	}

	ag = restoredAg.(*Reservation)
	return nil
}
