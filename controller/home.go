package controller

import (
	"net/http"

	"github.com/bonfy/go-mega-code/vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates[tpName].Execute(w, &v)
}

func check(username, password string) bool {
	if username == "bonfy" && password == "abc123" {
		return true
	}
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !check(username, password) {
			v.AddError("username password not correct, please input again")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
