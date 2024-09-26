package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/milo-sobral/PoolPlanner/internal/handlers"
)

type Template struct {
	tmpl *template.Template
}

func NewTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {

	// Create a simple server which serves index.html
	e := echo.New()
	e.Renderer = NewTemplate()

	// Routes
	e.GET("/", handlers.HandleDefault)
	e.GET("/language-eng", handlers.HandleLanguageEng)
	e.GET("/language-fra", handlers.HandleLanguageFra)

	// Start server
	e.Logger.Fatal(e.Start(":42069"))
}
