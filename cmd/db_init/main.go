package main

import (
	"log"

	"github.com/bonfy/go-mega-code/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "bonfy",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
		{
			Username:     "rene",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Posts: []model.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
