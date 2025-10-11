package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Load environment variables from a .env file and 
// Establish connection to the MongoDB database
func Connect() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error in reading .env file")
	}

	dbUrl := os.Getenv("MONGODB_URI")
	if dbUrl == "" {
		log.Fatal("Cannot find the database.")
	}

	fmt.Println("Database connection successful: ", dbUrl)

	clientOption := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(clientOption)

	if err != nil {
		return nil
	}
	return client
}

// Initializ a global MongoDB client
var Client *mongo.Client = Connect()

// Return a reference to the specified 
// collection from connected MongoDB database
func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error in reading .env file")
	}
	dbName:=os.Getenv("DATABASE_NAME")
	fmt.Println("Database name: ", dbName)
	collection := client.Database(dbName).Collection(collectionName)

	if collection == nil {
		return nil
	}
	return collection
}
