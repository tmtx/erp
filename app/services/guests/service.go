package guests

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/guests/http"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/validator"
)

type Service struct {
	app.BasicService
}

func New(basicService app.BasicService) Service {
	return Service{basicService}
}

func (s *Service) NewRouter() server.Router {
	return &http.Router{s.CommandBus}
}

func (s *Service) RegisterCommandCallbacks() {
	s.CommandBus.Subscribe(
		app.CreateGuest,
		func(p bus.MessageParams) (validator.Messages, error) {
			params := app.CreateGuestParams{
				Email: p["email"].(string),
				Name:  p["name"].(string),
			}
			return s.Create(params)
		},
	)
}
