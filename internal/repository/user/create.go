package user

import (
	"context"
	"database/sql"
	"errors"
	"test/internal/entity"
)

func (r *Repo) Create(ctx context.Context, user entity.User) error {
	q := "INSERT INTO users(login, password_hash, nickname, firstname, lastname) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	t := r.manager.GetTxOrDefault(ctx)
	var id int64
	err := t.QueryRowContext(ctx, q, user.Login, user.PasswordHash, user.Nickname, user.Firstname, user.Lastname).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UserLoginExists(ctx context.Context, login string) (bool, error) {
	q := `SELECT * FROM USERS WHERE login = $1`
	t := r.manager.GetTxOrDefault(ctx)
	var user User
	err := t.GetContext(ctx, &user, q, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Repo) UserNicknameExists(ctx context.Context, nickname string) (bool, error) {
	q := `SELECT * FROM USERS WHERE nickname = $1`
	t := r.manager.GetTxOrDefault(ctx)
	var user User
	err := t.GetContext(ctx, &user, q, nickname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
