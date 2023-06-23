package database

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() error {
	if client == nil {
		ctx := context.TODO()
		var err error
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
		if err != nil {
			return err
		}
	}
	return nil
}

func GetApplications() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Connect()

	cursor, err := client.Database("whereiapplied").Collection("applications").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []bson.M

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func AddApplication(name string, link string) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Connect()
	result, err := client.Database("whereiapplied").Collection("applications").InsertOne(ctx, bson.M{
		"name": name,
		"link": link,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteApplication(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.Database("whereiapplied").Collection("applications").DeleteOne(ctx, bson.M{
		"_id": id,
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func UpdateApplication(id string, name string, link string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.Database("whereiapplied").Collection("applications").UpdateOne(ctx, bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"name": name,
			"link": link,
		},
	})
	if err != nil {
		return result, err
	}
	return result, nil
}
