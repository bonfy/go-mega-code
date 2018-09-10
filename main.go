package main

import (
	"net/http"

	"github.com/bonfy/go-mega-code/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
