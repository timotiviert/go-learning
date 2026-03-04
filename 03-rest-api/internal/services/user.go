package services

import (
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
	return s.repo.Create(user)
}
