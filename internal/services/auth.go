package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hoainam183/todo-app/internal/models"
	"github.com/hoainam183/todo-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	if existing, err := s.repo.GetByUsername(username); err == nil && existing != nil && existing.ID != "" {
		return nil, errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		ID:       uuid.NewString(),
		Username: username,
		Password: string(hashed),
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	// Hide password before returning
	u.Password = ""
	return u, nil
}
