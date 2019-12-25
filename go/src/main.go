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
	e.GET("/hello", handler.MainPage())
	e.GET("/hello/get", handler.GetPage())
	e.POST("/hello/post", handler.PostPage())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
