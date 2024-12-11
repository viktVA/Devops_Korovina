package user

import (
	"context"
	"test/internal/entity"
)

func (s *Service) GetUser(ctx context.Context, userId entity.UserID) (entity.User, error) {
	return s.userRepo.GetById(ctx, userId)
}
