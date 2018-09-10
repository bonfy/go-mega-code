package controller

import (
	"net/http"

	"github.com/bonfy/go-mega-code/vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}
