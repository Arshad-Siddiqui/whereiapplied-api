package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Arshad-Siddiqui/whereiapplied-api/database"
	"github.com/Arshad-Siddiqui/whereiapplied-api/router"
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
	handler := router.New()
	port := getPort()

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
