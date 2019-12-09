package users

import (
	"encoding/gob"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/app/users/http"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
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
	// For session serialization
	gob.Register(User{})
	return &http.Router{s.CommandBus, s}
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
				[]string{"email"},
			)
			if err != nil {
				return validator.Messages{}, err
			}

			params := app.UpdateUserInfoParams{
				Email: p["email"].(string),
			}
			return s.UpdateUserInfo(params)
		},
	)
}
