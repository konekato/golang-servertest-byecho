package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"handler"
	"interceptor"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	// Echo instance
	e := echo.New()

	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", handler.MainPage())
	e.GET("/user/get/:id", handler.GetUser(), interceptor.BasicAuth())
	e.POST("/user/post", handler.PostUser())
	e.GET("/user/page1", handler.Template())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
