package garments

import (
	"context"

	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Garment struct {
	ID        string
	Title     string
	Category  string
	Color     string
	wearCount int
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

//GetAll - Get All garments for collection
func GetAll() ([]Garment, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	issues := []Garment{}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return issues, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return issues, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Garment{}
		err := cur.Decode(&t)
		if err != nil {
			return issues, err
		}
		issues = append(issues, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}
