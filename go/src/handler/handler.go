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

func DBIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		// mysqlへ接続
		db, err := sql.Open("mysql", "root@/go_db")
		log.Println("Connected to mysql.")

		// 接続でエラーが発生した場合の処理
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		ins, err := db.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(3, "l")

		return c.JSON(http.StatusOK, err)
	}
}

func DBOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		// mysqlへ接続
		db, err := sql.Open("mysql", "root@/go_db")
		log.Println("Connected to mysql.")

		// 接続でエラーが発生した場合の処理
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		u := new(User)

		// データの取得
		if err := db.QueryRow("SELECT * FROM users WHERE id = 3").Scan(&u.Email, &u.Name); err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, u)
	}
}

func DBUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// mysqlへ接続
		db, err := sql.Open("mysql", "root@/go_db")
		log.Println("Connected to mysql.")

		// 接続でエラーが発生した場合の処理
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		upd, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		upd.Exec("lll", 3)

		return c.JSON(http.StatusOK, err)
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

func PostForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "posttest", nil)
	}
}

func PostTest() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
