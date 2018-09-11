package vm

import (
	"log"

	"github.com/bonfy/go-mega-code/model"
)

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// LoginViewModelOp strutc
type LoginViewModelOp struct{}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

// CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}
