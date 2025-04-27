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
func (r *UserRepository) GetAll() (users []*models.User, err error) {
	result, err := r.collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}
	defer result.Close(context.TODO())
	if err = result.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
func (r *UserRepository) FindByID(id bson.ObjectID) (user *models.User, err error) {

	err = r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (user *models.User, err error) {

	err = r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) InsertOne(user *models.User) (bson.ObjectID, error) {

	result, err := r.collection.InsertOne(context.TODO(), user)

	return result.InsertedID.(bson.ObjectID), err
}

func (r *UserRepository) UpdateOne(id bson.ObjectID, newValues *models.UserUpdate) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": newValues}
	return r.collection.UpdateOne(context.TODO(), filter, update)
}

func (r *UserRepository) DeleteOne(id bson.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
