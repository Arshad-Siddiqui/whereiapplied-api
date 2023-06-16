package database

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	id, err := AddApplication("Google", "https://google.com")
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
	result, err := AddApplication("Google", "https://google.com")
	if err != nil {
		t.Error(err)
	}
	var id string
	if result != nil {
		id = getId(result)
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
	result, err := AddApplication("Google", "https://google.com")
	if err != nil {
		t.Error(err)
	}
	var id string
	if result != nil {
		id = getId(result)
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

	// Clear the database
	err = client.Database("whereiapplied").Collection("applications").Drop(context.TODO())
	if err != nil {
		panic("Error clearing the database")
	}
}

func getId(result *mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}
