package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a App) HeathCheck(c echo.Context) error {
	return c.String(http.StatusOK, "All systems operational!!")
}
