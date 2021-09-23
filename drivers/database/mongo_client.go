package database

import (
	"context"
	"crud/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient the specific client for mongo db
type MongoDBClient struct {
	connection *mongo.Database
}

func (client *MongoDBClient) connect(config config.Database) {
	// Set client options
	fmt.Printf("mongodb://%v:%v", config.Host, config.Port)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port))

	// Connect to MongoDB
	connection, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = connection.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	client.connection = connection.Database(config.Database)
}

// GetConnection retrieve database connection as singleton
func (client *MongoDBClient) GetConnection() *mongo.Database {
	if !client.isConnected() {
		client.connect(config.Config.MongoDB)
		fmt.Printf("Mongo db connected %v ", client.connection)
	}
	return client.connection
}

// isConnected return if there a database connection in memory
func (client MongoDBClient) isConnected() bool {
	return client.connection != nil
}
