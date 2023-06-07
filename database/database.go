package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	ctx := context.TODO()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func GetApplications() []bson.M {
	var results []bson.M

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("whereiapplied").Collection("applications")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
		fmt.Println(result)
	}
	log.Println(results)
	return results
}
