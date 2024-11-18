package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/milosobral/PoolPlanner/internal/handlers"
)

func NewTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("templates/*.html")),
	}
}

type Template struct {
	tmpl *template.Template
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

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}
