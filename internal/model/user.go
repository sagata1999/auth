package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUser struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int32
}

type UserInfo struct {
	Name  string
	Email string
	Role  int32
}
