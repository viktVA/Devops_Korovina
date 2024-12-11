package user

import (
	"context"
	"test/internal/entity"
	userService "test/internal/service/user"
)

type UserService interface {
	Create(ctx context.Context, user userService.CreateUser) error
	Auth(ctx context.Context, login string, password string) (token string, err error)
	GetUser(ctx context.Context, userId entity.UserID) (entity.User, error)
}
