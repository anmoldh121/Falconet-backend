package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	echo *echo.Echo
	db   *mongo.Client
}

func (s *Server) Listen(addr string) {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})
	s.echo.POST("/auth", s.Register)
	s.echo.Start(addr)
}

func New(client *mongo.Client) Server {
	e := echo.New()
	return Server{
		echo: e,
		db:   client,
	}
}
