package garments

import (
	"context"
	"fmt"

	"github.com/hirvoin/outfits-server/graph/model"
	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Garment struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Category   string             `bson:"category"`
	Color      string             `bson:"color"`
	WearCount  int                `bson:"wearCount"`
	IsFavorite bool               `bson:"isFavorited"`
	ImageUri   string             `bson:"imageUri"`
}

// Formats collection Garment to model Garment
func (dbGarment *Garment) FormatToModel() *model.Garment {
	var garment model.Garment
	garment.ID = dbGarment.ID.Hex()
	garment.Title = dbGarment.Title
	garment.Color = dbGarment.Color
	garment.Category = dbGarment.Category
	garment.IsFavorite = dbGarment.IsFavorite
	garment.WearCount = dbGarment.WearCount
	garment.ImageURI = dbGarment.ImageUri
	return &garment
}

// Insert new garment to the collection.
func CreateGarment(garment Garment) (*mongo.InsertOneResult, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := database.GetMongoClient()
	if err != nil {
		return nil, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	//Perform InsertOne operation & validate against the error.
	res, err := collection.InsertOne(context.TODO(), garment)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Edit garment by replacing garment with new data
func EditGarment(garment Garment) (Garment, error) {
	newGarment := Garment{}

	//Get MongoDB connection using connectionhelper.
	client, err := database.GetMongoClient()
	if err != nil {
		return newGarment, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	filter := bson.D{primitive.E{Key: "_id", Value: garment.ID}}

	//Perform FindOneAndReplace & validate against the error.
	err = collection.FindOneAndReplace(context.TODO(), filter, garment).Decode(&newGarment)

	if err != nil {
		return Garment{}, err
	}

	return newGarment, nil
}

// Get Garment by id from collection.
func GetGarmentById(id string) (Garment, error) {
	garment := Garment{}

	// Define filter query for fetching specific garment from collection
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return garment, err
	}

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
		return Garment{}, err
	}

	return garment, nil
}

// Get Garments by ids from collection.
func GetGarmentsByIds(ids []primitive.ObjectID) ([]Garment, error) {
	garments := []Garment{}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		return garments, err
	}

	//Define filter query for fetching specific documents from collection
	filter := bson.M{"_id": bson.M{"$in": ids}}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	// Perform FindOne operation & validate against the error.
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return garments, err
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

	return garments, nil
}

// Get all garments from collection.
func GetAll() ([]Garment, error) {
	garments := []Garment{}

	//Define filter query for fetching specific document from collection
	filter := bson.D{{}}

	//Get MongoDB connection.
	client, err := database.GetMongoClient()
	if err != nil {
		fmt.Println(err)
		return garments, err
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.GARMENTS)

	// Perform Find operation & validate against the error.
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return garments, err
	}

	// Map result to slice
	for cur.Next(context.TODO()) {
		garment := Garment{}
		err := cur.Decode(&garment)
		if err != nil {
			fmt.Println(err)
			return garments, err
		}
		garments = append(garments, garment)
	}

	// Once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(garments) == 0 {
		return garments, mongo.ErrNoDocuments
	}

	return garments, nil
}
