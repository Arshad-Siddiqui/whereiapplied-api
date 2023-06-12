package controller

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
)

func ListApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	applications, err := database.GetApplications()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(applications)
}
