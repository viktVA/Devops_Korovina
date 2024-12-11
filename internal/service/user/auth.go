package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"test/internal/entity"
	"test/pkg/errs"
	"time"
)

func (s *Service) Auth(ctx context.Context, login string, password string) (token string, err error) {
	user, err := s.userRepo.GetByLogin(ctx, login)
	if err != nil {
		err = errs.UserNotFound
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		err = errs.ErrInvalidPassword
		return
	}

	jwtToken := jwt.New(jwt.SigningMethodES256)

	jwtToken.Claims = entity.Claims{
		UserID:         user.ID,
		Login:          user.Login,
		Nickname:       user.Nickname,
		StandardClaims: jwt.StandardClaims{},
		ExpiresAt:      time.Now().Add(30 * time.Second),
	}

	token, err = jwtToken.SignedString(s.cfg.PrivateKey)
	return
}
