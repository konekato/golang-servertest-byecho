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
	e.GET("/user", handler.MainPage())
	e.GET("/user/get/dbin", handler.DBIn())
	e.GET("/user/get/dbout", handler.DBOut())
	e.GET("/user/get/dbupdate", handler.DBUpdate())
	e.GET("/user/get/:id", handler.GetUser(), interceptor.BasicAuth())
	e.POST("/user/post", handler.PostUser())
	e.GET("/user/page1", handler.Page1Template())
	e.GET("/user/page2", handler.Page2Template())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// package main

// import (
// 	"database/sql" //ここでパッケージをimport
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
// )

// //引っ張ってきたデータを割り当てるための構造体を用意
// type Person struct {
// 	ID   int
// 	Name string
// }

// func main() {

// 	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
// 	db, err := sql.Open("mysql", "root@/go_db")
// 	log.Println("Connected to mysql.")

// 	//接続でエラーが発生した場合の処理
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// //データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
// 	// rows, err := db.Query("SELECT * FROM users")
// 	// defer rows.Close()
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// //レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
// 	// for rows.Next() {
// 	// 	var person Person //構造体Person型の変数personを定義
// 	// 	err := rows.Scan(&person.ID, &person.Name)

// 	// 	if err != nil {
// 	// 		panic(err.Error())
// 	// 	}
// 	// 	fmt.Println(person.ID, person.Name)

// 	// }

// 	var person Person

// 	if err := db.QueryRow("SELECT * FROM users WHERE id = 3").Scan(&person.ID, &person.Name); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(person.ID, person.Name)
// }
