package users

import (
	"context"
	"fmt"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type User struct {
	Id             *event.UUID
	Email          string
	HashedPassword string
}

func (u *User) ApplyEvent(e event.Event) {
	switch e.Key {
	case app.UserCreated:
		params := e.Params

		if params["email"] != nil {
			u.Email = params["email"].(string)
		}
		if params["hashed_password"] != nil {
			u.HashedPassword = params["hashed_password"].(string)
		}
		u.Id = e.EntityId
	}
}

func (u *User) Validate() (bool, *validator.Messages) {
	return false, nil
}

func (s *Service) RestoreAggregateRootById(id *event.UUID) (app.Aggregate, error) {
	targetEvents := []bus.MessageKey{
		app.UserCreated,
	}

	events, err := s.EventRepository.FindAllWithFilter(
		context.Background(),
		event.Filter{
			"key":       event.Filter{"$in": targetEvents},
			"entity_id": id.String(),
		},
	)

	if err != nil {
		return nil, err
	}

	u := User{}
	for _, e := range events {
		u.ApplyEvent(e)
	}

	return app.Aggregate(&u), nil
}

func (s *Service) RestoreAggregateRootByEmail(email string) (app.Aggregate, error) {
	filter := event.Filter{
		"key":          app.UserCreated,
		"params.email": email,
	}

	// Find initial user creation event
	e, err := s.EventRepository.FindOneWithFilter(
		context.Background(),
		filter,
	)

	if err != nil {
		return nil, err
	}
	if e.EntityId == nil {
		return nil, fmt.Errorf("Entity id can not be nil")
	}

	return s.RestoreAggregateRootById(e.EntityId)
}
