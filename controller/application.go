package controller

import (
	"encoding/json"
	"io"
	"log"
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
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}
			defer r.Body.Close()

			var data map[string]string
			if err := json.Unmarshal(body, &data); err != nil {
				log.Println(err)
				return
			}
			name := data["name"]
			url := data["url"]
			result, err := database.AddApplication(name, url)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			id := result.InsertedID.(primitive.ObjectID).Hex()
			idData := []byte(`{"id": "` + id + `"}`)
			w.Write([]byte("TEST"))
			w.Write(idData)
		} else {
			// get form data
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}
			name := r.FormValue("name")
			url := r.FormValue("url")
			result, err := database.AddApplication(name, url)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			id := result.InsertedID.(primitive.ObjectID).Hex()
			idData := []byte(`{"id": "` + id + `"}`)
			w.Write(idData)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}
