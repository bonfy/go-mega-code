package main

import (
	"html/template"
	"net/http"
)

// User struct
type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "bonfy"}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &user)
	})
	http.ListenAndServe(":8888", nil)
}
