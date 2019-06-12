package driver

import (
	"team-members/utils"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB returns a connected *mongo.Client
func ConnectDB() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	utils.LogFatal(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	utils.LogFatal(err)

	fmt.Println("Connected to MongoDB!")

	return client
}

// DisConnectDB disconnects connected *mongo.Client
func DisConnectDB(c *mongo.Client) {
	err := c.Disconnect(context.TODO())

	utils.LogFatal(err)

	fmt.Println("Connection to MongoDB closed.")
}
