package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"test/internal/entity"
)

func (s *Service) Create(ctx context.Context, user CreateUser) error {
	ok, err := s.userRepo.UserLoginExists(ctx, user.Login)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("user login already существует")
	}
	ok, err = s.userRepo.UserNicknameExists(ctx, user.Nickname)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("user nickname already существует")
	}
	hash, err := passwordHash(user.Password)
	if err != nil {
		return err
	}
	err = s.userRepo.Create(ctx, entity.User{
		Login:        user.Login,
		PasswordHash: hash,
		Nickname:     user.Nickname,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
	})
	if err != nil {
		return err

	}
	return nil

}

func passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
