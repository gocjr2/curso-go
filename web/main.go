package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:abc123456@tcp(localhost:3306)/cursogoweb?charset=utf8")

func main() {

	stmt, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?,?);")
	checkErr(err)

	_, err = stmt.Exec("Post 2", "Content 2")
	checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		post := Post{Id: 1, Title: "Unamed Post", Body: "No content"}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println((http.ListenAndServe(":8080", nil)))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
