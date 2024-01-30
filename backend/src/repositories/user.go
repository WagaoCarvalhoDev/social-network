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
	insertUser := "insert into users (nameUser, nick, email, passwd) values(?, ?, ?, ?)"
	statement, err := u.db.Prepare(insertUser)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.NameUser, user.Nick, user.Email, user.Passwd)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil
}
