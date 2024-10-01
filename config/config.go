package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongoDB() (*mongo.Client, error) {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority")

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connect to MongoDB
    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    log.Println("Connected to MongoDB!")
    return client, nil
}
