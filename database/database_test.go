package database

import (
	"testing"
)

func TestAddApplication(t *testing.T) {
	// Connect to the database
	err := Connect()
	if err != nil {
		t.Error(err)
	}
	err = AddApplication("Google", "https://google.com")
	if err != nil {
		t.Error(err)
	}
}
