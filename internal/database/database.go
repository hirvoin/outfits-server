package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

// Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

// Constants just to hold required database config's.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "outfits-app"
	GARMENTS         = "garments"
	OUTFITS          = "outfits"
	USERS            = "users"
)

func GetMongoClient() (*mongo.Client, error) {
	// Perform connection creation operation only once.
	mongoOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
		mongoUser := os.Getenv("MONGO_DB_USER")
		mongoPassword := os.Getenv("MONGO_DB_PASSWORD")
		MONGO_DB_URI := "mongodb+srv://" + mongoUser + ":" + mongoPassword + "@aduzki.ifriy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

		// Set client options and connect to database
		clientOptions := options.Client().ApplyURI(MONGO_DB_URI)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
