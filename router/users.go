package router

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
)

func usersMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", controller.SignUp)
	mux.HandleFunc("/login", controller.Login)
	return mux
}
