package users

import (
	"context"
	"fmt"

	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

// Insert the user to collection
func (user *User) Create() {
	client, err := database.GetMongoClient()

	if err != nil {
		fmt.Println(err)
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.USERS)

	hashedPassword, err := HashPassword(user.Password)

	// Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), &User{ID: primitive.NewObjectID(), Username: user.Username, Password: hashedPassword})

	if err != nil {
		fmt.Println(err)
	}
}

// Hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Compares password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Check if a user exists in database by given username
func GetUserIdByUsername(username string) (primitive.ObjectID, error) {
	user := User{}

	client, err := database.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.USERS)

	// Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	return user.ID, nil
}

func (user *User) Authenticate() bool {
	dbUser := User{}

	client, err := database.GetMongoClient()
	if err != nil {
		fmt.Println(err)
		return false
	}

	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}

	collection := client.Database(database.DB).Collection(database.USERS)

	err = collection.FindOne(context.TODO(), filter).Decode(&dbUser)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return CheckPasswordHash(user.Password, dbUser.Password)
}
