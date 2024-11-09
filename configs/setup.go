package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this function first configures the client to use the correct uri and check for errors
// we definede a timeout of 10 seconds we wanted to use when trying to connect
// then we check if there is an error connecting to the database and cancel the connection if connection period exceeds 10s
// thn we pinged the database to test our conn and returned client instance
// createn a DB variable instance of the connectDB.
func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to MongoDB")
	return client
}

// client instance
var DB *mongo.Client = connectDB()

//getting the database collections

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
