package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitializeMongoDB(uri ...string) *mongo.Client {
	var defaultURI string
	if len(uri) == 0 {
		defaultURI = "mongodb://localhost:27017"
	} else {
		defaultURI = uri[0]
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(defaultURI))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}
