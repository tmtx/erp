package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
)

type Router struct {
	Service app.SpacesService
}

func (r *Router) GetPrefix() string {
	return "/spaces"
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/available", Callback: r.GetAvailable, Method: http.MethodGet},
	}
}

func (r *Router) GetAvailable(c echo.Context) error {
	availableSpaces, err := r.Service.GetAllAvailable()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, availableSpaces)
}
