package auth_test

import (
	"fmt"
	"testing"

	"github.com/Arshad-Siddiqui/whereiapplied-api/auth"
	"github.com/golang-jwt/jwt/v5"
)

func TestValidateJWT(t *testing.T) {
	id := "12345"
	tokenStr, err := auth.CreateJWT(id)
	if err != nil {
		t.Error("Error signing with key")
	}

	token, err := auth.ValidateJWT(tokenStr)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("claims not ok")
	}

	if claims["id"] != id {
		t.Error("Id is different")
	}
	fmt.Println(claims["id"])
}
