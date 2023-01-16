package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {

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
