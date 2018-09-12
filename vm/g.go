package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

// SetCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}
