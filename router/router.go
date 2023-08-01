package router

import (
	"net/http"

	"github.com/rs/cors"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/applications/", http.StripPrefix("/applications", appsMux()))
	mux.Handle("/users/", http.StripPrefix("/users", usersMux()))

	mux.HandleFunc("/", index)
	return cors.Default().Handler(mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
