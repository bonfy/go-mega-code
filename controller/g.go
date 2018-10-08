package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	flashName      string
	store          *sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = "go-mega"
	flashName = "go-flash"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
