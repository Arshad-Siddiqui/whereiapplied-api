package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/Arshad-Siddiqui/whereiapplied-api/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	applications, err := database.GetApplications()
	if errors.InternalServer(w, &err) != nil {
		return
	}
	w.Write(applications)
}

// TODO: Fix this, currently doesn't update the database
func AddApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var application database.Application
	err := json.NewDecoder(r.Body).Decode(&application)
	if errors.StatusBadRequest(w, &err) != nil {
		return
	}
	result, err := database.AddApplication(application)
	if errors.InternalServer(w, &err) != nil {
		return
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	w.Write([]byte(id))
}
