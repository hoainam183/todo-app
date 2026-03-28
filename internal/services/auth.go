package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/hoainam183/todo-app/internal/models"
	"github.com/hoainam183/todo-app/internal/repository"
	"github.com/hoainam183/todo-app/pkg/common/apperrors"
)

type AuthService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
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

type LoginResult struct {
	UserId string
	Token  string
}

func (s *AuthService) Login(username, password string) (*LoginResult, error) {
	if username == "" || password == "" {
		return &LoginResult{UserId: "", Token: ""}, errors.New("username and password are required")
	}

	// Get user from database
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return &LoginResult{UserId: "", Token: ""}, errors.New("invalid username or password")
	}

	// Verify password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return &LoginResult{UserId: "", Token: ""}, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := s.generateJWT(user)
	if err != nil {
		return &LoginResult{UserId: "", Token: ""}, errors.New("failed to generate token")
	}

	return &LoginResult{UserId: user.ID, Token: token}, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
