package database

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUser(t *testing.T) {
	setup()
	user := User{
		Email:    "testemail",
		Password: "testpassword",
	}
	result, err := AddUser(user)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("Expected result to be returned")
	}
}

func TestFindUser(t *testing.T) {
	setup()
	user := User{
		Email:    "testemail",
		Password: "testpassword",
	}

	result, err := AddUser(user)
	if err != nil {
		t.Error(err)
	}

	insertedId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		t.Log("failed to turn the result id into a string")
	}

	insertedIdString := insertedId.Hex()

	userFromDb, err := FindUser(insertedIdString)
	if err != nil {
		t.Log(err)
	}

	if !reflect.DeepEqual(user, userFromDb) {
		t.Log("Epic fail. Users are not the same.")
	}
}
