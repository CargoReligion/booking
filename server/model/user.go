package model

import (
	"github.com/google/uuid"
)

type UserRole string

const (
	RoleCoach   UserRole = "coach"
	RoleStudent UserRole = "student"
)

type User struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber"`
	Role        UserRole  `db:"user_role" json:"role"`
}
