package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	}
}

func PostUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	}
}
