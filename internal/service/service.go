package service

import (
	"context"

	"github.com/sagata1999/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.CreateUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
