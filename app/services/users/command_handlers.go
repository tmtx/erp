package users

import (
	"context"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

func (s *Service) Login(p app.LoginUserParams) (validator.Messages, error) {
	ag := &aggregates.User{
		Email: p.Email,
		Base:  aggregates.Base{Repository: s.EventRepository},
	}

	restoredAggregate, err := aggregates.RestoreFromEmail(ag, ag.Email)
	if err != nil {
		return nil, err
	}
	ag = restoredAggregate.(*aggregates.User)

	if isValid, validatorMessages := ValidateLoginUser(ag, p); !isValid {
		return validatorMessages, nil
	}

	params := bus.MessageParams{
		"email": p.Email,
	}
	e := event.New(app.UserLoggedIn, params, nil)
	return nil, s.EventRepository.Store(
		context.Background(),
		e,
	)
}

func (s *Service) UpdateUserInfo(p app.UpdateUserInfoParams) (validator.Messages, error) {
	if isValid, validatorMessages := ValidateUserInfo(p); !isValid {
		return validatorMessages, nil
	}

	params := bus.MessageParams{
		"email": p.Email,
		"id":    p.UserId,
	}
	e := event.New(app.UserInfoUpdated, params, p.UserId)
	return nil, s.EventRepository.Store(
		context.Background(),
		e,
	)
}
