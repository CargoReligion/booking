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
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name"`
	Email string    `db:"email" json:"email"`
	Role  UserRole  `db:"user_role" json:"role"`
}
