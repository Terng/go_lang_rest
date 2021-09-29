package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// timeout after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection
func getConnect() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

// Create creating a task in a mongo or documents db
func Create(task *Task) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnect()
	defer cancel()
	defer client.Disconnect(ctx)
	task.ID = primitive.NewObjectID()

	result, err := client.Database("task").Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		log.Printf("Could not create Task: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}
