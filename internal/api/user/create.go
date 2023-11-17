package user

import (
	"context"
	"log"

	"github.com/sagata1999/auth/internal/converter"
	desc "github.com/sagata1999/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("[api] create user request: %#v", req.GetUser())
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetUser()))
	if err != nil {
		log.Printf("[api] error %v", err)
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
