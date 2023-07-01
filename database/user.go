package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(user User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.Database("whereiapplied").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
