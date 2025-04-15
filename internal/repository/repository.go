package repository

import (
	"context"

	"github.com/sagata1999/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, create_model *model.CreateUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
