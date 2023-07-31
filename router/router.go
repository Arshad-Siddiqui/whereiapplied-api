package router

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
	"github.com/rs/cors"
)

func New() http.Handler {
	mux := http.NewServeMux()

	appsMux := http.NewServeMux()
	usersMux := http.NewServeMux()
	appsMux.HandleFunc("/", controller.ListApplications)
	appsMux.HandleFunc("/add", controller.AddApplication)
	usersMux.HandleFunc("/users/signup", controller.SignUp)
	usersMux.HandleFunc("/users/login", controller.Login)

	mux.Handle("/applications/", http.StripPrefix("/applications", appsMux))
	mux.Handle("/users/", http.StripPrefix("/users", usersMux))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	return cors.Default().Handler(mux)
}
