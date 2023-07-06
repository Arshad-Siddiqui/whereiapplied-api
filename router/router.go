package router

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
	"github.com/rs/cors"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/applications", controller.ListApplications)
	mux.HandleFunc("/applications/add", controller.AddApplication)
	mux.HandleFunc("/users/add", controller.AddUser)
	mux.HandleFunc("/users/login", controller.Login)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	return cors.Default().Handler(mux)
}
