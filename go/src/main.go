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
	e.GET("/", handler.MainPage())
	e.GET("/user/get/dbin", handler.DBIn())
	e.GET("/user/get/dbout", handler.DBOut())
	e.GET("/user/get/dbupdate", handler.DBUpdate())
	e.GET("/user/get/:id", handler.GetUser(), interceptor.BasicAuth())
	e.POST("/user/post", handler.PostUser())
	e.GET("/user/page1", handler.Page1Template())
	e.GET("/user/page2", handler.Page2Template())
	e.GET("/user", handler.PostForm())
	e.POST("user/posttest", handler.PostTest())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
