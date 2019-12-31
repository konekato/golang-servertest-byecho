package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", handler.MainPage())
	e.GET("/user/get/:id", handler.GetUser())
	e.POST("/user/post", handler.PostUser())
	e.GET("/user/page1", handler.Page1Template())
	e.GET("/user/page2", handler.Page2Template())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
