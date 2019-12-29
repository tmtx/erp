package reservations

import (
	"time"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/reservations/http"
	"github.com/tmtx/res-sys/pkg/bus"
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
	return &http.Router{s.CommandBus, s}
}

// TODO: maybe parameter parsing should be done somewhere else
func (s *Service) RegisterCommandCallbacks() {
	s.CommandBus.Subscribe(
		app.CreateReservation,
		func(p bus.MessageParams) (validator.Messages, error) {
			if _, ok := p["spaceId"].(float64); !ok {
				return nil, app.TypedError{Type: app.ParamTypeError}
			}
			spaceId := uint(p["spaceId"].(float64))

			startDate, err := time.Parse(
				time.RFC3339,
				p["startDate"].(string),
			)
			if err != nil {
				return nil, err
			}

			endDate, err := time.Parse(
				time.RFC3339,
				p["endDate"].(string),
			)
			if err != nil {
				return nil, err
			}

			params := app.CreateReservationParams{
				GuestEmail: p["email"].(string),
				GuestName:  p["name"].(string),
				SpaceId:    spaceId,
				StartDate:  startDate,
				EndDate:    endDate,
			}
			return s.Create(params)
		},
	)
}
