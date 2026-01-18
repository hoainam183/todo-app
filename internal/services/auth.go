package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/hoainam183/todo-app/internal/models"
	"github.com/hoainam183/todo-app/internal/repository"
	"github.com/hoainam183/todo-app/pkg/common/apperrors"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username, password string) error {
	if username == "" || password == "" {
		return apperrors.ErrMissingField
	}

	if existing, err := s.repo.GetByUsername(username); err == nil && existing != nil && existing.ID != "" {
		return apperrors.ErrUserExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &models.User{
		ID:       uuid.NewString(),
		Username: username,
		Password: string(hashed),
	}

	if err := s.repo.Create(u); err != nil {
		return err
	}
	return nil
}
