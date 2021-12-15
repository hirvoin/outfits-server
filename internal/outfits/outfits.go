package outfits

import (
	"context"

	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Outfit struct {
	ID       primitive.ObjectID   `bson:"_id"`
	Date     primitive.DateTime   `bson:"date"`
	Garments []primitive.ObjectID `bson:"garments"`
}

// Insert a new document in the collection.
func CreateOutfit(outfit Outfit) (*mongo.InsertOneResult, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := database.GetMongoClient()
	if err != nil {
		return nil, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.OUTFITS)

	//Perform InsertOne operation & validate against the error.
	res, err := collection.InsertOne(context.TODO(), outfit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Outfit by id from collection.
func GetOutfitById(id string) (Outfit, error) {
	outfit := Outfit{}

	//Define filter query for fetching specific document from collection
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objId}}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return outfit, err
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.OUTFITS)

	// Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&outfit)
	if err != nil {
		return outfit, err
	}

	return outfit, nil
}

// Get all outfits from collection.
func GetAll() ([]Outfit, error) {
	outfits := []Outfit{}

	//Define filter query for fetching specific document from collection
	filter := bson.D{{}}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return outfits, err
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.OUTFITS)

	// Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return outfits, findError
	}

	// Map result to slice
	for cur.Next(context.TODO()) {
		outfit := Outfit{}
		err := cur.Decode(&outfit)
		if err != nil {
			return outfits, err
		}
		outfits = append(outfits, outfit)
	}

	// Once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(outfits) == 0 {
		return outfits, mongo.ErrNoDocuments
	}

	return outfits, nil
}
