package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Arshad-Siddiqui/whereiapplied-api/controller"
	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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
	port := getPort()
	handler := cors.Default().Handler(mux)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
