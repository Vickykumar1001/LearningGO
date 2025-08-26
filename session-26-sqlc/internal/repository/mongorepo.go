package repository

import (
	"context"
	"log"
	"session-23-gin-jwt/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepo struct {
	db *mongo.Client
}

func (m MongoRepo) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	coll := m.db.Database("users").Collection("user")
	log.Println("User here ----", user)
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func (m MongoRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	coll := m.db.Database("users").Collection("user")
	// Creates a query filter to match documents in which the "username" is

	filter := bson.D{{"username", userName}}
	// Retrieves the first matching document
	var result models.User
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m MongoRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	coll := m.db.Database("users").Collection("user")
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (m MongoRepo) UpdateUser(ctx context.Context, id interface{}, user models.User) error {
	coll := m.db.Database("users").Collection("user")
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{
		{"username", user.Username},
		{"firstname", user.FirstName},
		{"secondname", user.SeconName},
	}}}

	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}

func (m MongoRepo) DeleteUser(ctx context.Context, id interface{}) error {
	coll := m.db.Database("users").Collection("user")
	filter := bson.D{{"id", id}}
	_, err := coll.DeleteOne(ctx, filter)
	return err
}

func NewMongoRepo(db *mongo.Client) DbRepository {
	return &MongoRepo{db}

}
