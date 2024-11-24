package service

import (
	"github.com/cargoreligion/booking/server/model"
	"github.com/cargoreligion/booking/server/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}
