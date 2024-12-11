package entity

import "time"

type User struct {
	ID           UserID
	Login        string
	PasswordHash string
	Nickname     string
	Firstname    string
	Lastname     string
	CreatedAt    time.Time
	DeletedAt    time.Time
}
