package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetId(result *mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func GeneralContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
