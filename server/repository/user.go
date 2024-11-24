package repository

import (
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/model"
)

type UserRepository struct {
	dbc db.DbClient
}

func NewUserRepository(dbc db.DbClient) *UserRepository {
	return &UserRepository{dbc: dbc}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	query := `SELECT id, name, email, phone_number, role FROM users`
	err := r.dbc.Select(&users, query)
	return users, err
}
