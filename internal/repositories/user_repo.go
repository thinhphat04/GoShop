package repositories

import (
	"API/internal/models"
	"API/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser saves a new user to the database
func (r *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetUsers retrieves all users from the database
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}
