package router

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
)

func appsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.ListApplications)
	mux.HandleFunc("/add", controller.AddApplication)
	return mux
}
