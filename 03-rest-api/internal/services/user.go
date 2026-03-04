package services

import (
	"github.com/timotiviert/go-learning/03-rest-api/internal/auth"
	"github.com/timotiviert/go-learning/03-rest-api/internal/database"
	"github.com/timotiviert/go-learning/03-rest-api/internal/models"
)

type UserService struct {
	repo database.UserRepository
}

func New(r database.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Create(user *models.RegisterUsers) (*models.User, error) {
	// Hash password using bcrypt.
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(user.Email, user.Username, hashedPassword)
}
