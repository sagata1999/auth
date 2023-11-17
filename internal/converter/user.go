package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sagata1999/auth/internal/model"
	desc "github.com/sagata1999/auth/pkg/user_v1"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  desc.Role(info.Role),
	}
}

func ToUserInfoFromDesc(user *desc.CreateUser) *model.CreateUser {
	return &model.CreateUser{
		Name:            user.Name,
		Email:           user.Email,
		Role:            int32(user.GetRole()),
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
	}
}
