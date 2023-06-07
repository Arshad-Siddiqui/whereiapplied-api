package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/joho/godotenv"
)

func listApplications(w http.ResponseWriter, r *http.Request) {
	applications := database.GetApplications()
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(applications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.Connect()
	mux := http.NewServeMux()
	mux.HandleFunc("/applications", listApplications)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
