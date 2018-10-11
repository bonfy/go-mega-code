package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bonfy/go-mega-code/vm"
	"github.com/gorilla/mux"
)

type home struct{}

func (h home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unFollowHandler))
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/explore", middleAuth(exploreHandler))
	r.HandleFunc("/reset_password_request", resetPasswordRequestHandler)
	r.HandleFunc("/reset_password/{token}", resetPasswordHandler)
	r.HandleFunc("/", middleAuth(indexHandler))

	http.Handle("/", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	page := getPage(r)
	username, _ := getSessionUser(r)
	if r.Method == http.MethodGet {
		flash := getFlash(w, r)
		v := vop.GetVM(username, flash, page, pageLimit)
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		body := r.Form.Get("body")
		errMessage := checkLen("Post", body, 1, 180)
		if errMessage != "" {
			setFlash(w, r, errMessage)
		} else {
			err := vm.CreatePost(username, body)
			if err != nil {
				log.Println("add Post error:", err)
				w.Write([]byte("Error insert Post in database"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
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

		errs := checkLogin(username, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "register.html"
	vop := vm.RegisterViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	page := getPage(r)
	vop := vm.ProfileViewModelOp{}
	v, err := vop.GetVM(sUser, pUser, page, pageLimit)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) does not exist", pUser)
		w.Write([]byte(msg))
		return
	}
	templates[tpName].Execute(w, &v)
}

func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile_edit.html"
	username, _ := getSessionUser(r)
	vop := vm.ProfileEditViewModelOp{}
	v := vop.GetVM(username)
	if r.Method == http.MethodGet {
		err := templates[tpName].Execute(w, &v)
		if err != nil {
			log.Println(err)
		}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		aboutme := r.Form.Get("aboutme")
		log.Println(aboutme)
		if err := vm.UpdateAboutMe(username, aboutme); err != nil {
			log.Println("update Aboutme error:", err)
			w.Write([]byte("Error update aboutme"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusSeeOther)
	}
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)

	err := vm.Follow(sUser, pUser)
	if err != nil {
		log.Println("Follow error:", err)
		w.Write([]byte("Error in Follow"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

func unFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)

	err := vm.UnFollow(sUser, pUser)
	if err != nil {
		log.Println("UnFollow error:", err)
		w.Write([]byte("Error in UnFollow"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "explore.html"
	vop := vm.ExploreViewModelOp{}
	username, _ := getSessionUser(r)
	page := getPage(r)
	v := vop.GetVM(username, page, pageLimit)
	templates[tpName].Execute(w, &v)
}

func resetPasswordRequestHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "reset_password_request.html"
	vop := vm.ResetPasswordRequestViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")

		errs := checkResetPasswordRequest(email)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)

		} else {
			log.Println("Send mail to", email)
			vopEmail := vm.EmailViewModelOp{}
			vEmail := vopEmail.GetVM(email)
			var contentByte bytes.Buffer
			tpl, _ := template.ParseFiles("templates/email.html")

			if err := tpl.Execute(&contentByte, &vEmail); err != nil {
				log.Println("Get Parse Template:", err)
				w.Write([]byte("Error send email"))
				return
			}
			content := contentByte.String()
			go sendEmail(email, "Reset Password", content)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	username, err := vm.CheckToken(token)
	if err != nil {
		w.Write([]byte("The token is no longer valid, please go to the login page."))
	}

	tpName := "reset_password.html"
	vop := vm.ResetPasswordViewModelOp{}
	v := vop.GetVM(token)

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		log.Println("Reset password for ", username)
		r.ParseForm()
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkResetPassword(pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := vm.ResetUserPassword(username, pwd1); err != nil {
				log.Println("reset User password error:", err)
				w.Write([]byte("Error update user password in database"))
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
