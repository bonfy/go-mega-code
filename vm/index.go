package vm

import "github.com/bonfy/go-mega-code/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := model.User{Username: "bonfy"}
	u2 := model.User{Username: "rene"}

	posts := []model.Post{
		model.Post{User: u1, Body: "Beautiful day in Portland!"},
		model.Post{User: u2, Body: "The Avengers movie was so cool!"},
	}

	vm := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return vm
}
