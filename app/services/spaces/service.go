package spaces

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/app/services/spaces/http"
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
