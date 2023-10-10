package database

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Arshad-Siddiqui/whereiapplied-api/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Application struct {
	Name    string `json:"name"`
	Applied bool   `json:"applied"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	Website string `json:"website"`
}

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
	ctx, cancel := util.GeneralContext()
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

func AddApplication(app Application) (*mongo.InsertOneResult, error) {
	ctx, cancel := util.GeneralContext()
	defer cancel()
	result, err := client.Database("whereiapplied").Collection("applications").InsertOne(ctx, bson.M{
		"name":    app.Name,
		"applied": app.Applied,
		"status":  app.Status,
		"date":    app.Date,
		"website": app.Website,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteApplication(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := util.GeneralContext()
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
	ctx, cancel := util.GeneralContext()
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

func GetAppCount() (int64, error) {
	ctx, cancel := util.GeneralContext()
	defer cancel()

	count, err := client.Database("whereiapplied").Collection("applications").CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
