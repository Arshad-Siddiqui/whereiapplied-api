package auth_test

import (
	"testing"

	"github.com/Arshad-Siddiqui/whereiapplied-api/auth"
)

func TestGetClaims(t *testing.T) {
	idArr := []string{"12345", "helloworld", "fooBar"}

	for _, id := range idArr {
		tokenStr, err := auth.CreateJWT(id)
		if err != nil {
			t.Error(err, "Error signing with key")
		}

		claims, err := auth.GetClaims(tokenStr)
		if err != nil {
			t.Error(err, "Error generating claims")
		}
		if claims["id"] != id {
			t.Error("Id is different")
		}
	}
}
