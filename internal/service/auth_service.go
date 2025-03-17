package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ruanv123/api-go-crud/internal/model"
	"github.com/ruanv123/api-go-crud/internal/repository"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(user *model.User) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !repository.ComparePassword(user.Password, password) {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Name,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("minha-chave-secreta"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authService) Register(user *model.User) error {
	return s.userRepo.CreateUser(user)
}
