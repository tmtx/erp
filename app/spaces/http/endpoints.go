package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
)

type Router struct {
	Service app.SpacesService
}

func (r *Router) GetPrefix() string {
	return "/spaces"
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/available", Callback: r.GetAvailableEndpoint, Method: http.MethodGet},
	}
}

func (r *Router) GetAvailableEndpoint(c echo.Context) error {
	availableSpaces, err := r.Service.GetAllAvailable()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, availableSpaces)
}
