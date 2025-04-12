package service

import (
	"week_9_crud/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
}

type UserService struct {
	r UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{r: r}
}

func (s UserService) GetAllUsers() ([]models.User, error) {
	return s.r.GetAllUsers()
}
