package server

import (
	"net/http"

	"github.com/chatApp/controller"
	"github.com/labstack/echo/v4"
)

func (s *Server) Register(c echo.Context) error {
	resp, err := controller.Register(c)
	if err != nil {
		return echo.NewHTTPError(500, "Error in server")
	}
	return c.JSON(http.StatusOK, resp)
}
