package vm

import "github.com/bonfy/go-mega-code/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string) IndexViewModel {
	u1, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts}
	v.SetCurrentUser(username)
	return v
}
