package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/pkg/bus"
)

type Router struct {
	CommandBus bus.MessageBus
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/create", Callback: r.CreateEndpoint, Method: http.MethodGet},
	}
}

func (r *Router) CreateEndpoint(c echo.Context) error {
	params := app.CreateGuestParams{
		Name:  "Test",
		Email: "kf@karlis.dev",
	}
	r.CommandBus.Dispatch(bus.NewCommand(app.CreateGuest, params))
	return c.String(http.StatusOK, "Hello, World!")
}
