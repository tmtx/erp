package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerFunc echo.HandlerFunc

type RouteHandler struct {
	Path     string
	Callback HandlerFunc
	Method   string
}

type Router interface {
	Routes() []RouteHandler
}

type Server struct {
	echo *echo.Echo
}

func New(routers []Router) Server {
	s := Server{
		echo.New(),
	}
	for _, r := range routers {
		for _, rh := range r.Routes() {
			s.RegisterRouteHandler(rh)
		}
	}

	return s
}

func (s *Server) Start(conn string) {
	s.echo.Logger.Fatal(s.echo.Start(conn))
}

func (s *Server) RegisterRouteHandler(r RouteHandler) {
	switch r.Method {
	case http.MethodPost:
		s.echo.POST(r.Path, echo.HandlerFunc(r.Callback))
	case http.MethodGet:
		fallthrough
	default:
		s.echo.GET(r.Path, echo.HandlerFunc(r.Callback))
	}
}
