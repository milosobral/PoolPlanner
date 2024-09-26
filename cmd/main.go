package main

import (
	"github.com/labstack/echo/v4"
	"github.com/milo-sobral/PoolPlanner/internal/handlers"
)

func main() {

	// Create a simple server which serves index.html
	e := echo.New()
	e.Renderer = newTemplate()

	// Routes
	e.GET("/", handlers.HandleHeader)

	// Start server
	e.Logger.Fatal(e.Start(":42069"))

}
