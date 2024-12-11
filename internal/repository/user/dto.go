package user

import (
	"database/sql"
	"test/internal/entity"
)

type User struct {
	ID           entity.UserID  `db:"id"`
	Login        string         `db:"login"`
	PasswordHash string         `db:"password_hash"`
	Nickname     string         `db:"nickname"`
	FirstName    sql.NullString `db:"firstname"`
	LastName     sql.NullString `db:"lastname"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	DeletedAt    sql.NullTime   `db:"deleted_at"`
}

type LoginUser struct {
	Login    string
	Password string
}
type CreateUser struct {
	Login     string
	Hash      string
	Nickname  string
	Firstname string
	Lastname  string
}

func (u *User) toUser() entity.User {
	return entity.User{
		ID:           u.ID,
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
		Nickname:     u.Nickname,
		Firstname:    u.FirstName.String,
		Lastname:     u.LastName.String,
		CreatedAt:    u.CreatedAt.Time,
		DeletedAt:    u.DeletedAt.Time,
	}
}
