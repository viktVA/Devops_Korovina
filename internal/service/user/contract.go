package user

import (
	"context"
	"test/internal/entity"
)

type UserRepo interface {
	Create(ctx context.Context, user entity.User) error
	UserLoginExists(ctx context.Context, login string) (bool, error)
	UserNicknameExists(ctx context.Context, nickname string) (bool, error)
	GetByLogin(ctx context.Context, login string) (entity.User, error)
	GetById(ctx context.Context, id entity.UserID) (entity.User, error)
}
