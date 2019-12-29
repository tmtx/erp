package spaces

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/aggregates"
)

func (s *Service) GetAllAvailable() ([]app.Aggregate, error) {
	// TODO: implement
	spaces := []app.Aggregate{
		&aggregates.Space{Id: 1},
		&aggregates.Space{Id: 3},
		&aggregates.Space{Id: 93},
		&aggregates.Space{Id: 42},
		&aggregates.Space{Id: 371},
	}

	return spaces, nil
}
