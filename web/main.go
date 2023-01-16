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
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println((http.ListenAndServe(":8080", nil)))
}
