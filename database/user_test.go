package database

import (
	"testing"
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
