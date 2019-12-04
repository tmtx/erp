package users

import (
	"context"
	"fmt"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

func (s *Service) Login(p app.LoginUserParams) (*validator.Messages, error) {
	u, err := s.RestoreAggregateRootByEmail(p.Email)
	fmt.Println(u)
	if err != nil {
		return nil, err
	}

	if isValid, validatorMessages := ValidateLoginUser(u, p); !isValid {
		return validatorMessages, fmt.Errorf("Validation failed")
	}

	return nil, s.EventRepository.Store(
		context.Background(),
		event.New(app.UserLoggedIn, p),
	)
}

func (s *Service) RestoreAggregateRootById(id event.UUID) (*User, error) {
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
		u.ApplyEvent(&e)
	}

	return &u, nil
}

func (s *Service) RestoreAggregateRootByEmail(email string) (*User, error) {
	filter := event.Filter{
		"key":          app.UserCreated,
		"params.email": email,
	}

	// Find initial user creation event
	e := event.New(app.UserCreated, app.CreateUserParams{})
	err := s.EventRepository.FindOneWithFilter(
		context.Background(),
		filter,
		&e,
	)
	if err != nil {
		return nil, err
	}
	if e.EntityId == nil {
		return nil, fmt.Errorf("Entity id can not be nil")
	}

	return s.RestoreAggregateRootById(*e.EntityId)
}
