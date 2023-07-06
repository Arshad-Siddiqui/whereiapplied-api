package database

import (
	"errors"

	"github.com/Arshad-Siddiqui/whereiapplied-api/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(user User) (*mongo.InsertOneResult, error) {
	ctx, cancel := util.DbContext()
	defer cancel()

	// TODO: Check if this even works
	if CheckUserExists(user.Email) {
		return nil, errors.New("user already exists")
	}

	result, err := client.Database("whereiapplied").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindUser(id string) (User, error) {
	ctx, cancel := util.DbContext()
	defer cancel()

	db := client.Database("whereiapplied")
	collection := db.Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	filter := bson.M{"_id": objectID}

	var user User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func CheckUserExists(email string) bool {
	ctx, cancel := util.DbContext()
	defer cancel()

	db := client.Database("whereiapplied")
	collection := db.Collection("users")

	filter := bson.M{"email": email}

	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	return err == nil
}

func CheckLogin(possibleUser User) (User, error) {
	ctx, cancel := util.DbContext()
	defer cancel()

	db := client.Database("whereiapplied")
	collection := db.Collection("users")

	filter := bson.M{"email": possibleUser.Email, "password": possibleUser.Password}

	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}
