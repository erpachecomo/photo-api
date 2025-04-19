package repository

import (
	"context"

	"github.com/erpachecomo/photo-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) FindByID(id bson.ObjectID) (user *models.User, err error) {

	err = r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) InsertOne(user *models.User) (bson.ObjectID, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)
	return result.InsertedID.(bson.ObjectID), err
}
