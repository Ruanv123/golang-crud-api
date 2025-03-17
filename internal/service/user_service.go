package service

import (
	"errors"

	"github.com/ruanv123/api-go-crud/internal/model"
	"github.com/ruanv123/api-go-crud/internal/repository"
)

type UserService interface {
	Profile(id int) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) Profile(id int) (*model.User, error) {
	user, err := u.userRepo.GetByID(id)

	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}
