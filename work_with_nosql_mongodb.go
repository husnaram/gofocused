package gofocused

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TheStudent struct {
	Name  string `bson:"name,omitempty"`
	Grade int    `bson:"grade,omitempty"`
}

var ctx = context.Background()

var credential = options.Credential{
	Username: "mandasia",
	Password: "sastra@dee@",
}

func TestConnected() {
	clientOpt := options.Client().SetAuth(credential)
	clientOpt.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOpt)
	if err != nil {
		return
	}

	fmt.Println(ctx)
	fmt.Println(client)
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("learn_golang"), nil
}

func Insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, TheStudent{"Wick", 3})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, TheStudent{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}
