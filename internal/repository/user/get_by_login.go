package user

import (
	"context"
	"test/internal/entity"
)

func (r *Repo) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	q := `select * from users WHERE login = $1`
	t := r.manager.GetTxOrDefault(ctx)
	var user User

	err := t.GetContext(ctx, &user, q, login)

	if err != nil {
		return entity.User{}, err
	}

	return user.toUser(), err

}
