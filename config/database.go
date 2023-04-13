package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envs Config = GetEnv()

func ConnectDB() *mongo.Client {
  uri := envs.Database.GetUri()
  clientOptions := options.Client().ApplyURI(uri)
  ctx := context.Background()

  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Database connection stablished")

  return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
  dbName := envs.Database.Name
  collection := client.Database(dbName).Collection(collectionName)

  return collection
}
