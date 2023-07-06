package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/auth"
	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if handleStatusBadRequest(w, &err) != nil {
		return
	}

	result, err := database.AddUser(user)
	if handleInternalServerError(w, &err) != nil {
		return
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	w.Write([]byte(id))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if handleStatusBadRequest(w, &err) != nil {
		return
	}

	user, err = database.CheckLogin(user)
	if handleInternalServerError(w, &err) != nil {
		return
	}

	jwt, err := auth.CreateJWT(user.ID)
	if handleInternalServerError(w, &err) != nil {
		return
	}

	w.Write([]byte(jwt))
}

func handleInternalServerError(w http.ResponseWriter, err *error) error {
	if *err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte((*err).Error()))
		return *err
	}
	return nil
}

func handleStatusBadRequest(w http.ResponseWriter, err *error) error {
	if *err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte((*err).Error()))
		return *err
	}
	return nil
}
