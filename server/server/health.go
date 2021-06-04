package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) health(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
