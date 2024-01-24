package repositories

import (
	"backend/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) CreateUser(user models.User) (uint64, error) {
	return 0, nil
}
