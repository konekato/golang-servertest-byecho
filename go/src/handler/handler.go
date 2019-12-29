package handler

import (
	"database/sql"
	"net/http"

	"log"

	_ "github.com/go-sql-driver/mysql"
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
	Name  string `query:"name"`
	Email string `query:"id"`
}

func DBOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mysqlへ接続
		db, err := sql.Open("mysql", "root@/go_db")
		log.Println("Connected to mysql.")

		//接続でエラーが発生した場合の処理
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		u := new(User)

		if err := db.QueryRow("SELECT * FROM users WHERE id = 1").Scan(&u.Email, &u.Name); err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, u)
	}
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

func Template() echo.HandlerFunc {
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
