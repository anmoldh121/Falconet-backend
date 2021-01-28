package server

import (
	"encoding/json"
	"net/http"

	"github.com/chatApp/controller"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetPeer(c echo.Context) error {
	json_map := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return echo.NewHTTPError(500, "Error in server")
	}
	resp, err := controller.GetPeer(json_map["Ph"], s.db)
	if err != nil {
		return echo.NewHTTPError(500, "Error in server")
	}
	return c.JSON(http.StatusOK, resp)
}
