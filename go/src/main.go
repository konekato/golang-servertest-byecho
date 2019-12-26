package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", handler.MainPage())
	e.GET("/user/get/:id", handler.GetUser())
	e.POST("/user/post", handler.PostUser())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
