package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
