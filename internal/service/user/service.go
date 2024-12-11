package user

import "test/internal/cfg"

type Service struct {
	userRepo UserRepo
	cfg      *cfg.Config
}

func NewService(userRepo UserRepo, cfg *cfg.Config) *Service {
	return &Service{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

type CreateUser struct {
	Login     string
	Password  string
	Nickname  string
	Firstname string
	Lastname  string
}
