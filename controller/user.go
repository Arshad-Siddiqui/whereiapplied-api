package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/auth"
	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/Arshad-Siddiqui/whereiapplied-api/errors"
	"github.com/Arshad-Siddiqui/whereiapplied-api/util"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if errors.StatusBadRequest(w, &err) != nil {
		return
	}

	result, err := database.AddUser(user)
	if errors.InternalServer(w, &err) != nil {
		return
	}
	id := util.GetId(result)
	w.Write([]byte(id))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if errors.StatusBadRequest(w, &err) != nil {
		return
	}

	user, err = database.CheckLogin(user)
	if errors.InternalServer(w, &err) != nil {
		return
	}

	jwt, err := auth.CreateJWT(user.ID)
	if errors.InternalServer(w, &err) != nil {
		return
	}

	w.Write([]byte(jwt))
}
