package services

import (
	"github.com/erpachecomo/photo-api/internal/models"
	errors "github.com/erpachecomo/photo-api/pkg/utils"
	"github.com/erpachecomo/photo-api/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(db *mongo.Database) *UserService {
	userRepository := repository.NewUserRepository(db)
	return &UserService{
		repository: userRepository,
	}

}

func (s *UserService) GetUser(id string) (*models.User, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, mongo.ErrNoDocuments
	}

	return s.repository.FindByID(oid)
}

func (s *UserService) PostUser(user *models.User) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.repository.FindByEmail(user.Email)

	if err == nil && existingUser != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	result, err := s.repository.InsertOne(user)
	if err != nil {
		return nil, err
	}
	newUser, err := s.repository.FindByID(result)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

// Update the user and returns the updated user
func (s *UserService) PutUser(id string, user *models.UserUpdate) (*models.User, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, mongo.ErrNoDocuments
	}

	res, err := s.repository.UpdateOne(oid, user)

	if res.ModifiedCount == 0 || err != nil {
		return nil, err
	}

	updatedUser, err := s.repository.FindByID(oid)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil

}
