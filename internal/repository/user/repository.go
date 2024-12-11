package user

import (
	"test/internal/txmanager"
)

type Repo struct {
	manager *txmanager.TxManager
}

func New(manager *txmanager.TxManager) *Repo {
	return &Repo{
		manager: manager,
	}
}
