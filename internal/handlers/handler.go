package handlers

import (
	"github.com/labstack/echo/v4"
)

func HandleHeader(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}
