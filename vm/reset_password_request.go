package vm

import (
	"log"

	"github.com/bonfy/go-mega-code/model"
)

// ResetPasswordRequestViewModel struct
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

// ResetPasswordRequestViewModelOp struct
type ResetPasswordRequestViewModelOp struct{}

// GetVM func
func (ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

// CheckEmailExist func
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("Can not find email:", email)
		return false
	}
	return true
}
