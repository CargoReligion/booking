package repository

import (
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	dbc db.DbClient
}

func NewUserRepository(dbc db.DbClient) *UserRepository {
	return &UserRepository{dbc: dbc}
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, phone_number, user_role FROM stepful_user WHERE id = $1`
	err := r.dbc.GetSingleEntity(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	query := `SELECT id, name, phone_number, user_role FROM stepful_user`
	err := r.dbc.Select(&users, query)
	return users, err
}
