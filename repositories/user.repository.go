package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/joshnavdev/loans-server/config"
	"github.com/joshnavdev/loans-server/models"

	// "github.com/joshnavdev/loans-server/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "user")

type UserRepository struct {
  collection *mongo.Collection
}

func NewUserRepository() UserRepository {
  return UserRepository{
    collection: userCollection,
  }
}

func (repo *UserRepository) Create(user models.User) error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  result, err := repo.collection.InsertOne(ctx, user)

  fmt.Println(result)

  if err != nil {
    return err
  }

  return nil
}

func (repo *UserRepository) CountByEmail(email string) (int64, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  userCount, err := repo.collection.CountDocuments(ctx, bson.M{"email": email})

  return userCount, err
}
