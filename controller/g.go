package controller

import "html/template"

var (
	homeController home
	templates      map[string]*template.Template
)

func init() {
	templates = PopulateTemplates()
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
