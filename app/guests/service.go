package guests

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/guests/http"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
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
