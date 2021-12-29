package users

import (
	"context"

	"github.com/hirvoin/outfits-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"log"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func (user *User) Create() {
	client, err := database.GetMongoClient()

	if err != nil {
		log.Fatal(err)
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.USERS)

	hashedPassword, err := HashPassword(user.Password)

	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), &User{ID: primitive.NewObjectID(), Username: user.Username, Password: hashedPassword})

	if err != nil {
		log.Fatal(err)
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (primitive.ObjectID, error) {
	user := User{}

	client, err := database.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.USERS)

	// Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user.ID, nil
}

func (user *User) Authenticate() bool {
	foundUser := User{}

	client, err := database.GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return false
	}

	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}

	collection := client.Database(database.DB).Collection(database.USERS)

	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return CheckPasswordHash(user.Password, foundUser.Password)
}
