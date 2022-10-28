package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllPosts(coll *mongo.Collection) []bson.M {
	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func AddNewPost(coll *mongo.Collection, title string, text string) error {
	if title == "" || text == "" {
		return errors.New("Validation failed: title or text is empty")
	}

	coll.InsertOne(context.TODO(), bson.D{{Key: "title", Value: title}, {Key: "text", Value: text}})
	return nil
}
