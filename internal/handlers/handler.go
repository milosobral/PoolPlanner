package handlers

import (
	"github.com/labstack/echo/v4"
)

func HandleHeader(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}

func HandleLanguageFra(c echo.Context) error {
	return c.Render(200, "explanation-fra", nil)
}

func HandleLanguageEng(c echo.Context) error {
	return c.Render(200, "explanation-eng", nil)
}
