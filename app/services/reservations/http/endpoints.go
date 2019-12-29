package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/pkg/bus"
)

type Router struct {
	CommandBus bus.MessageBus
	Service    app.ReservationsService
}

func (r *Router) GetPrefix() string {
	return "/reservations"
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/create", Callback: r.Create, Method: http.MethodPost},
		{Path: "/list", Callback: r.GetAll, Method: http.MethodGet},
	}
}

func (r *Router) Create(c echo.Context) error {
	params, err := server.MessageParamsFromJSONBody(c)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	validatorMessages, err := r.CommandBus.DispatchSync(
		bus.NewCommand(app.CreateReservation, params),
	)
	if err != nil {
		if typedError, ok := err.(app.TypedError); ok {
			if typedError.Type == app.UnableToRestoreAggregate {
				return c.NoContent(http.StatusBadRequest)
			} else {
				return c.NoContent(http.StatusInternalServerError)
			}
		}
	}

	if len(validatorMessages) > 0 {
		return server.ValidationErrorResponse(c, validatorMessages)
	}

	return server.SuccessResponse(c)
}

func (r *Router) GetAll(c echo.Context) error {
	availableSpaces, err := r.Service.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, availableSpaces)
}
