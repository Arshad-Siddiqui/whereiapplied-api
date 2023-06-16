package database

import (
	"context"
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
	setup()
	err := AddApplication("Google", "https://google.com")
	if err != nil {
		t.Error(err)
	}
}

func TestGetApplications(t *testing.T) {
	setup()
	_, err := GetApplications()
	if err != nil {
		t.Error(err)
	}
}

func setup() {
	// Load the .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to the database
	err = Connect()
	if err != nil {
		panic("Error connecting to the database")
	}

	// Clear the database
	err = client.Database("whereiapplied").Collection("applications").Drop(context.TODO())
	if err != nil {
		panic("Error clearing the database")
	}
}
