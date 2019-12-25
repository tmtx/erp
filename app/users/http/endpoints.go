package http

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/pkg/bus"
)

type Router struct {
	CommandBus bus.MessageBus
	Service    app.UserService
}

func (r *Router) GetPrefix() string {
	return "/users"
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/login", Callback: r.Login, Method: http.MethodPost, Public: true},
		{Path: "/update", Callback: r.UpdateUserInfoParams, Method: http.MethodPost},
		{Path: "/me", Callback: r.CurrentUserInfo, Method: http.MethodGet},
	}
}

func (r *Router) Login(c echo.Context) error {
	params, err := server.MessageParamsFromJSONBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	validatorMessages, err := r.CommandBus.DispatchSync(bus.NewCommand(app.LoginUser, params))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	if len(validatorMessages) > 0 {
		return server.ValidationErrorResponse(c, validatorMessages)
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	user, err := r.Service.RestoreAggregateRootByEmail(params["email"].(string))
	if sess.Values["user"] == nil && user != nil {
		sess.Values["user"] = r.Service.Session(user)
		sess.Save(c.Request(), c.Response())
	}

	return server.SuccessResponse(c)
}

func (r *Router) UpdateUserInfoParams(c echo.Context) error {
	params, err := server.MessageParamsFromJSONBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	validatorMessages, err := r.CommandBus.DispatchSync(
		bus.NewCommand(app.UpdateUserInfo, params),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	if len(validatorMessages) > 0 {
		return server.ValidationErrorResponse(c, validatorMessages)
	}

	return server.SuccessResponse(c)
}

func (r *Router) CurrentUserInfo(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	userSession := sess.Values["user"]

	if userSession == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, userSession)

}
