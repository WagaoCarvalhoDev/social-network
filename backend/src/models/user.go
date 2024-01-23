package models

import "time"

type User struct {
	Id       uint64    `json:"id,omitempty"`
	NameUser string    `json:"nameUser,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Passwd   string    `json:"passwd,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}
