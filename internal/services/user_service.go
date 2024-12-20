package services

import (
	"API/internal/models"
	"API/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

// GetUsers retrieves all users
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetUsers()
}
