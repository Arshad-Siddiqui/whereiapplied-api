package database

import (
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

	result, err := client.Database("whereiapplied").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindUser(id string) (User, error) {
	// Queries the DB and then returns the ID and maybe an error.
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
