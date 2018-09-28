package vm

import "github.com/bonfy/go-mega-code/model"

// ProfileEditViewModel struct
type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

// ProfileEditViewModelOp struct
type ProfileEditViewModelOp struct{}

// GetVM func
func (ProfileEditViewModelOp) GetVM(username string) ProfileEditViewModel {

	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

// UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	return model.UpdateAboutMe(username, text)
}
