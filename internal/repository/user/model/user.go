package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type CreateUser struct {
	Name            string `db:"name"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	PasswordConfirm string `db:"password_confirm"`
	Role            int32  `db:"role"`
}

type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  int32  `db:"role"`
}
