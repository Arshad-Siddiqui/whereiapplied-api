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

func TestCheckUserExists(t *testing.T) {
	users := []User{
		{
			Email:    "testemail",
			Password: "testpassword",
		},
		{
			Email:    "testemail2",
			Password: "testpassword2",
		},
	}

	for _, user := range users {
		setup()

		_, err := AddUser(user)
		if err != nil {
			t.Error(err)
		}

		userExists := CheckUserExists(user.Email)
		if !userExists {
			t.Log("User should exist but doesn't")
		}
	}

	userExists := CheckUserExists("nonexistentemail")
	if userExists {
		t.Log("User should not exist but does")
	}
}

func TestCheckLogin(t *testing.T) {
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

	userFromDb, err := CheckLogin(user)
	if err != nil {
		t.Error(err)
	}

	if userFromDb.Email != user.Email || userFromDb.Password != user.Password {
		t.Log("Epic fail. Users are not the same.")
	}

	user2 := User{
		Email:    "testemail2",
		Password: "testpassword2",
	}

	_, err = CheckLogin(user2)
	if err == nil {
		t.Log("User should not exist but does")
	}
}
