package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/milo-sobral/PoolPlanner/internal/services/scraping"
)

const url string = "https://montreal.ca/lieux?mtl_content.lieux.installation.code=PISI&mtl_content.lieux.available_activities.code=ACT0"

type Data struct {
	Language LanguageInfo
	Pools    []scraping.Pool
}

func NewData() *Data {
	return &Data{
		Language: newLanguage("fra"),
		Pools:    scraping.GetPoolList(url),
	}
}

func HandleDefault(c echo.Context) error {
	data := NewData()
	return c.Render(200, "index.html", data)
}
