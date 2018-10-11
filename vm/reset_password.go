package vm

import (
	"github.com/bonfy/go-mega-code/model"
)

// ResetPasswordViewModel struct
type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

// ResetPasswordViewModelOp struct
type ResetPasswordViewModelOp struct{}

// GetVM func
func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

// CheckToken func
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

// ResetUserPassword func
func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
