package handlers

import (
	"github.com/labstack/echo/v4"
)

type Language struct {
	Language string
	FraClass string
	EngClass string
}

func newLanguage(lang string) Language {
	if lang == "fra" {
		return Language{Language: "fra", FraClass: "border-2 border-blue-500", EngClass: ""}
	} else {
		return Language{Language: "eng", EngClass: "border-2 border-blue-500", FraClass: ""}
	}
}

func HandleLanguageDefault(c echo.Context) error {
	data := newLanguage("fra")
	return c.Render(200, "index.html", data)
}

func HandleLanguageFra(c echo.Context) error {
	data := newLanguage("fra")
	return c.Render(200, "header", data)
}

func HandleLanguageEng(c echo.Context) error {
	data := newLanguage("eng")
	return c.Render(200, "header", data)
}
