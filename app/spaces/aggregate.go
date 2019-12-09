package spaces

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type Space struct {
	Id int `json:"id"`
}

func (s *Service) GetAllAvailable() ([]app.Aggregate, error) {
	// TODO: implement
	spaces := []app.Aggregate{
		&Space{Id: 1},
		&Space{Id: 3},
		&Space{Id: 93},
		&Space{Id: 42},
		&Space{Id: 371},
	}

	return []app.Aggregate(spaces), nil
}

func (s *Space) ApplyEvent(e event.Event) {
	// TODO: implement
}

func (s *Space) Validate() (bool, *validator.Messages) {
	// TODO: implement
	return false, nil
}
