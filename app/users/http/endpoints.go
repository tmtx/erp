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
		{Path: "/login", Callback: r.LoginEndpoint, Method: http.MethodGet},
	}
}

func (r *Router) LoginEndpoint(c echo.Context) error {
	params := app.LoginUserParams{
		Email: "kf@karlis.dev",
		// testtest
		HashedPassword: "$2a$07$R6im9feS1MoqIsRxysOVYuvcneRH9h1f7AUzf9WJEpbO.8vHaFSWC",
	}
	r.CommandBus.Dispatch(bus.NewCommand(app.LoginUser, params))
	return c.String(http.StatusOK, "Hello login")
}
