package database

import (
	"context"
	"testing"

	"github.com/Arshad-Siddiqui/whereiapplied-api/util"
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

var google = Application{
	Name:    "Google",
	Applied: true,
	Status:  "Pending",
	Date:    "2020-01-01",
	Website: "https://google.com",
}

func TestAddApplication(t *testing.T) {
	setup()
	id, err := AddApplication(google)
	if err != nil {
		t.Error(err)
	}
	if id == nil {
		t.Error("Expected id to be returned")
	}
}

func TestGetApplications(t *testing.T) {
	setup()
	_, err := GetApplications()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteApplication(t *testing.T) {
	setup()
	result, err := AddApplication(google)
	if err != nil {
		t.Error(err)
	}
	var id string
	if result != nil {
		id = util.GetId(result)
	} else {
		t.Error("Expected id to be returned")
	}
	delResult, err := DeleteApplication(id)
	if err != nil {
		t.Error(err)
	}
	if delResult == nil {
		t.Error("Expected id to be returned")
	}
}

func TestUpdateApplication(t *testing.T) {
	setup()
	result, err := AddApplication(google)
	if err != nil {
		t.Error(err)
	}
	var id string
	if result != nil {
		id = util.GetId(result)
	} else {
		t.Error("Expected id to be returned")
	}
	updateResult, err := UpdateApplication(id, "GoogleUpdated", "https://google.com/updated")
	if err != nil {
		t.Error(err)
	}
	if updateResult == nil {
		t.Error("Expected id to be returned")
	}
}

// <------ Helper functions ------>

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

	clearDB("applications")
	clearDB("users")
}

func clearDB(collection string) {
	err := client.Database("whereiapplied").Collection(collection).Drop(context.TODO())
	if err != nil {
		panic("Error clearing the database")
	}
}
