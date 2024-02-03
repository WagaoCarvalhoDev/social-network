package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	Id       uint64    `json:"id,omitempty"`
	NameUser string    `json:"nameUser,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Passwd   string    `json:"passwd,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

func (u *User) PrepareString() error {
	if erro := u.validate(); erro != nil {
		return erro
	}

	u.formatString()

	return nil
}

func (u *User) validate() error {
	if u.NameUser == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if u.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if u.Passwd == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *User) formatString() {
	u.NameUser = strings.TrimSpace(u.NameUser)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
