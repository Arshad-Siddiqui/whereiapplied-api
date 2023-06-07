package main

import (
	"log"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/joho/godotenv"
)

func listApplications(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.Connect()
	mux := http.NewServeMux()
	mux.HandleFunc("/applications", listApplications)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening on port 8080")
	server.ListenAndServe()
}
