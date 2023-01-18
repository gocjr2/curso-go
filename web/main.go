package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	Id      int
	Title   string
	Body    string
	Created string
}

var db, err = sql.Open("mysql", "root:abc123456@tcp(localhost:3306)/cursogoweb?charset=utf8")

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/js").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("js/"))))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	fmt.Println((http.ListenAndServe(":8080", r)))

	db.Close()
}

func listPosts() []Post {

	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body, &post.Created)
		items = append(items, post)
	}

	return items
}

func getPostById(id string) *Post {
	row := db.QueryRow("SELECT * FROM posts WHERE id=?", id)
	post := new(Post)
	row.Scan(&post.Id, &post.Title, &post.Body, &post.Created)
	return post
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout/default.html", "templates/list.html"))
	if err := t.ExecuteTemplate(w, "default.html", listPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	t := template.Must(template.ParseFiles("templates/layout/default.html", "templates/view.html"))
	if err := t.ExecuteTemplate(w, "default.html", getPostById(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func add() {
//stmt, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?,?);")
//checkErr(err)

//_, err = stmt.Exec("Post 2", "Content 2")
//checkErr(err)
//}
