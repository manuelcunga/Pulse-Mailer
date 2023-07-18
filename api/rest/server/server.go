package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Start() {

	server.Router.Logger.Fatal(server.Router.Start(":4000"))
}
