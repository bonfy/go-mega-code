package vm

import (
	"log"

	"github.com/bonfy/go-mega-code/model"
)

// RegisterViewModel struct
type RegisterViewModel struct {
	LoginViewModel
}

// RegisterViewModelOp struct
type RegisterViewModelOp struct{}

// GetVM func
func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

// CheckUserExist func
func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		return false
	}
	return true
}

// AddUser func
func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
