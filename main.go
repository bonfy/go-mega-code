package main

import (
	"net/http"

	"github.com/bonfy/go-mega-code/controller"
	"github.com/bonfy/go-mega-code/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", nil)
}
