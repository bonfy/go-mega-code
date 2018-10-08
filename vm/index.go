package vm

import "github.com/bonfy/go-mega-code/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string, flash string) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts, _ := u.FollowingPosts()
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts, flash}
	v.SetCurrentUser(username)
	return v
}

// CreatePost func
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
