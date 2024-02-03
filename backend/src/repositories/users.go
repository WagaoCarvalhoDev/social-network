package repositories

import (
	"backend/src/models"
	"database/sql"
	"fmt"
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

func (u users) FindByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nomeOuNick%

	rows, err := u.db.Query(
		"select id, nameUser, nick, email, CreateAt from users where nameUser LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.NameUser,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
