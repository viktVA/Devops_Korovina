package user

import (
	"context"
	"test/internal/entity"
)

func (r *Repo) GetById(ctx context.Context, id entity.UserID) (entity.User, error) {
	q := `select * from users WHERE id = $1`
	t := r.manager.GetTxOrDefault(ctx)
	var user User
	err := t.GetContext(ctx, &user, q, id)
	if err != nil {
		return entity.User{}, err
	}

	return user.toUser(), err

}
