package main

import (
	"html/template"
	"net/http"
)

// User struct
type User struct {
	Username string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User  User
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "bonfy"}
		v := IndexViewModel{Title: "Homepage", User: user}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
