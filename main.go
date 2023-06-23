package main

import (
	"log"
	"net/http"
	"os"

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
	port := getPort()
	corsMux := applyCORS(mux)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, corsMux))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func applyCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Update with your allowed origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Update with your allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Update with your allowed headers

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
