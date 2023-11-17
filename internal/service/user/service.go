package user

import (
	"github.com/sagata1999/auth/internal/client/db"
	"github.com/sagata1999/auth/internal/repository"
	"github.com/sagata1999/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
