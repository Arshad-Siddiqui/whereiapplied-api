package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// TODO: Fix this, currently doesn't update the database
func AddApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var application database.Application
	err := json.NewDecoder(r.Body).Decode(&application)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	result, err := database.AddApplication(application)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	w.Write([]byte(id))
}
