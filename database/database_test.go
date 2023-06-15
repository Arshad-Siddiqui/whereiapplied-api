package database

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	// Load the .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		t.Error("Error loading .env file")
	}

	// Connect to the database
	err = Connect()
	if err != nil {
		t.Error("Error connecting to the database")
	}
}
func TestAddApplication(t *testing.T) {
	// Load the .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		t.Error("Error loading .env file")
	}

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

func TestGetApplications(t *testing.T) {
	// Load the .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		t.Error("Error loading .env file")
	}

	// Connect to the database
	err = Connect()
	if err != nil {
		t.Error("Error connecting to the database")
	}
	_, err = GetApplications()
	if err != nil {
		t.Error(err)
	}
}
