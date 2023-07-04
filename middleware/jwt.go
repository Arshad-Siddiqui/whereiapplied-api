package middleware

import (
	"net/http"

	"github.com/Arshad-Siddiqui/whereiapplied-api/auth"
)

func extractToken(r *http.Request) string {
	// Extract token form header request.
	return r.Header.Get("Authorization")
}

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header or query string
		tokenString := extractToken(r)

		// Validate and parse the token claims
		claims, err := auth.GetClaims(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		id := claims["id"]

		// TODO: Add database check.
		if id == nil {
			http.Error(w, "No Id", http.StatusUnauthorized)
		}
		// Perform authorization logic based on token claims
		// For example, check if the user has necessary permissions

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
