package router

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
	"github.com/rs/cors"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/applications/", http.StripPrefix("/applications", appsMux()))
	mux.Handle("/users/", http.StripPrefix("/users", usersMux()))

	mux.HandleFunc("/", index)
	return cors.Default().Handler(mux)
}

func appsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.ListApplications)
	mux.HandleFunc("/add", controller.AddApplication)
	return mux
}

func usersMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", controller.SignUp)
	mux.HandleFunc("/login", controller.Login)
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
