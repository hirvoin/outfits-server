package garments

import (
	"context"
	"fmt"

	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Garment struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Category  string             `bson:"category"`
	Color     string             `bson:"color"`
	wearCount int                `bson:"wear_count"`
}

//CreateGarment - Insert a new document in the collection.
func CreateGarment(garment Garment) error {
	//Get MongoDB connection using connectionhelper.
	client, err := database.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), garment)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

// GetGarmentById - Get garment by id for collection
func GetGarmentById(id string) (Garment, error) {
	garment := Garment{}

	//Define filter query for fetching specific document from collection
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objId}}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return garment, err
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	// Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&garment)
	if err != nil {
		return garment, err
	}
	// Return result
	return garment, nil
}

// GetAll - Get All garments f collection
func GetAll() ([]Garment, error) {
	garments := []Garment{}

	//Define filter query for fetching specific document from collection
	filter := bson.D{{}}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return garments, err
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	// Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return garments, findError
	}

	// Map result to slice
	for cur.Next(context.TODO()) {
		garment := Garment{}
		err := cur.Decode(&garment)
		if err != nil {
			return garments, err
		}
		garments = append(garments, garment)
	}

	// Once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(garments) == 0 {
		return garments, mongo.ErrNoDocuments
	}

	fmt.Println(garments)
	// Return result
	return garments, nil
}
