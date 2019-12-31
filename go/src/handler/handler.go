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

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func PostUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	}
}

type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo{
	"サイトのタイトル",
}

func Page1Template() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートに渡す値

		data := struct {
			ServiceInfo
			Content_a string
			Content_b string
			Content_c string
			Content_d string
		}{
			ServiceInfo: serviceInfo,
			Content_a:   "雨が降っています。",
			Content_b:   "明日も雨でしょうか。",
			Content_c:   "台風が近づいています。",
			Content_d:   "Jun/11/2018",
		}
		return c.Render(http.StatusOK, "page1", data)
	}
}

func Page2Template() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートに渡す値

		data := struct {
			ServiceInfo
			Content_a string
			Content_b string
			Content_c string
			Content_d string
		}{
			ServiceInfo: serviceInfo,
			Content_a:   "晴れています。",
			Content_b:   "明日も雨でしょうか。",
			Content_c:   "台風が近づいています。",
			Content_d:   "Jun/12/2018",
		}
		return c.Render(http.StatusOK, "page2", data)
	}
}
