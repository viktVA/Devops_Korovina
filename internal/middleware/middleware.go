package middleware

import "test/internal/cfg"

type MwManager struct {
	cfg *cfg.Config
}

func New(cfg *cfg.Config) *MwManager {
	return &MwManager{
		cfg: cfg,
	}
}
