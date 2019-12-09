package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
)

type HandlerFunc echo.HandlerFunc

type RouteHandler struct {
	Path     string
	Callback HandlerFunc
	Method   string
	Public   bool
}

type Router interface {
	Routes() []RouteHandler
	GetPrefix() string
}

type Server struct {
	echo *echo.Echo
}

// Process is the middleware function.
func checkSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess == nil || sess.Values["user"] == nil {
			response := struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{"error", "Forbidden"}
			return c.JSON(http.StatusForbidden, response)
		}

		return next(c)
	}
}

func New(routers []Router) Server {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("AKssi29hha!o"))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	s := Server{
		e,
	}
	for _, r := range routers {
		for _, rh := range r.Routes() {
			s.RegisterRouteHandler(rh, r.GetPrefix())
		}
	}

	return s
}

func (s *Server) Start(conn string) {
	s.echo.Logger.Fatal(s.echo.Start(conn))
}

func (s *Server) RegisterRouteHandler(r RouteHandler, prefix string) {
	path := prefix + r.Path
	handlerFunc := echo.HandlerFunc(r.Callback)

	if !r.Public {
		handlerFunc = checkSession(handlerFunc)
	}

	switch r.Method {
	case http.MethodPost:
		s.echo.POST(path, handlerFunc)
	case http.MethodGet:
		fallthrough
	default:
		s.echo.GET(path, handlerFunc)
	}
}

func SuccessResponse(c echo.Context) error {
	response := struct {
		Status string `json:"status"`
	}{"ok"}
	return c.JSON(http.StatusOK, response)
}

func ValidationErrorResponse(c echo.Context, messages validator.Messages) error {
	formattedMessages := map[string]string{}
	for key, messageGroup := range messages {
		formattedMessages[key] = ""
		for _, m := range messageGroup {
			formattedMessages[key] += string(m) + "\n"
		}
	}
	response := struct {
		Status string            `json:"status"`
		Errors map[string]string `json:"errors"`
	}{"error", formattedMessages}
	return c.JSON(http.StatusOK, response)
}

func MessageParamsFromJSONBody(c echo.Context) (bus.MessageParams, error) {
	m := bus.MessageParams{}
	if err := c.Bind(&m); err != nil {
		return nil, err
	}
	return m, nil
}
