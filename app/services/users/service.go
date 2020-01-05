package users

import (
	"github.com/google/uuid"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/users/http"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

type Service struct {
	app.BasicService
}

func New(basicService app.BasicService) Service {
	return Service{
		basicService,
	}
}

func (s *Service) NewRouter() server.Router {
	return &http.Router{s.CommandBus, s, s.EventRepository}
}

func (s *Service) RegisterCommandCallbacks() {
	s.CommandBus.Subscribe(
		app.LoginUser,
		func(p bus.MessageParams) (validator.Messages, error) {
			err := app.CheckRequiredMessageParams(
				p,
				[]string{"email", "password"},
			)
			if err != nil {
				return validator.Messages{}, err
			}

			params := app.LoginUserParams{
				Email:    p["email"].(string),
				Password: p["password"].(string),
			}
			return s.Login(params)
		},
	)
	s.CommandBus.Subscribe(
		app.UpdateUserInfo,
		func(p bus.MessageParams) (validator.Messages, error) {
			err := app.CheckRequiredMessageParams(
				p,
				[]string{"email", "id"},
			)
			if err != nil {
				return validator.Messages{}, err
			}

			var userId uuid.UUID
			if id, ok := p["id"].(string); ok {
				userId, err = uuid.Parse(id)
				if err != nil {
					return nil, err
				}
			}

			params := app.UpdateUserInfoParams{
				Email:  p["email"].(string),
				UserId: &event.UUID{UUID: userId},
			}
			return s.UpdateUserInfo(params)
		},
	)
}

func (s *Service) Session(sessValues interface{}) server.Session {
	if u, ok := sessValues.(*aggregates.User); ok {
		return server.Session{
			Id:    u.Id.String(),
			Email: u.Email,
		}
	}

	return server.Session{}
}
