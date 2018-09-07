package main

import (
	"html/template"
	"net/http"
)

// User struct
type User struct {
	Username string
}

// Post struct
type Post struct {
	User User
	Body string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "bonfy"}
		u2 := User{Username: "rene"}

		posts := []Post{
			Post{User: u1, Body: "Beautiful day in Portland!"},
			Post{User: u2, Body: "The Avengers movie was so cool!"},
		}

		v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
