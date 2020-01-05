package aggregates

import (
	"context"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/pkg/event"
)

type Base struct {
	Id                *event.UUID      `json:"id"`
	Repository        event.Repository `json:"-"`
	Events            []event.Event
	AppliedEventCount uint
}

func (ag *Base) GetId() *event.UUID {
	return ag.Id
}

func (ag *Base) GetRepository() event.Repository {
	return ag.Repository
}

func (ag *Base) SetEvents(events []event.Event) {
	ag.Events = events
}

func (ag *Base) GetEvents() []event.Event {
	return ag.Events
}

func RestoreWithFilter(ag app.Aggregate, filter event.Filter) (app.Aggregate, error) {
	events, err := ag.GetRepository().FindAllWithFilter(
		context.Background(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	ag.SetEvents(events)

	for _, e := range events {
		ag.ApplyEvent(e)
	}

	return ag, nil
}

func RestoreFromId(ag app.Aggregate) (app.Aggregate, error) {
	filter := event.Filter{
		"key":       event.Filter{"$in": ag.GetTargetEvents()},
		"entity_id": ag.GetId().String(),
	}

	return RestoreWithFilter(ag, filter)
}

func RestoreFromEmail(
	ag app.Aggregate,
	email string,
) (app.Aggregate, error) {
	filter := event.Filter{
		"key":          event.Filter{"$in": ag.GetTargetEvents()},
		"params.email": email,
	}

	restoredAggregate, err := RestoreWithFilter(ag, filter)
	if err != nil {
		return nil, err
	}

	if restoredAggregate.GetId() != nil {
		return RestoreFromId(restoredAggregate)
	}

	return restoredAggregate, err
}
