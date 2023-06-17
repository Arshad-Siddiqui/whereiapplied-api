package main

import (
	"log"
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = database.Connect()
	if err != nil {
		panic(err)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/applications", controller.ListApplications)
	mux.HandleFunc("/applications/add", controller.AddApplication)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
