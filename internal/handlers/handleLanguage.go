package handlers

import (
	"github.com/labstack/echo/v4"
)

type LanguageInfo struct {
	Language string
	FraClass string
	EngClass string
}

func newLanguage(lang string) LanguageInfo {
	if lang == "fra" {
		return LanguageInfo{Language: "fra", FraClass: "border-2 border-blue-500", EngClass: ""}
	} else {
		return LanguageInfo{Language: "eng", EngClass: "border-2 border-blue-500", FraClass: ""}
	}
}

func HandleLanguageFra(c echo.Context) error {
	data := newLanguage("fra")
	return c.Render(200, "header", data)
}

func HandleLanguageEng(c echo.Context) error {
	data := newLanguage("eng")
	return c.Render(200, "header", data)
}
