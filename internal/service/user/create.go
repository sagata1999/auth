package user

import (
	"context"
	"log"

	"github.com/sagata1999/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, create_model *model.CreateUser) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, create_model)
		if errTx != nil {
			log.Printf("[service] error %v", errTx)
			return errTx
		}

		_, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
