package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAddApplication(t *testing.T) {
	// Load the .env file
	err := godotenv.Load("./.env.test")
	if err != nil {
		t.Error("Error loading .env file")
	}

	t.Log("MONGO_URI: ", os.Getenv("MONGO_URI:"))

	// Connect to the database
	err = Connect()
	if err != nil {
		t.Error("Error connecting to the database")
	}
	err = AddApplication("Google", "https://google.com")
	if err != nil {
		t.Error(err)
	}
}
