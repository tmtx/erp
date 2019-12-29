package spaces

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/spaces/http"
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
	return &http.Router{s}
}
