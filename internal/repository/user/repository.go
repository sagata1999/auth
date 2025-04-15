package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/sagata1999/auth/internal/client/db"
	"github.com/sagata1999/auth/internal/model"
	"github.com/sagata1999/auth/internal/repository"
	"github.com/sagata1999/auth/internal/repository/user/converter"
	modelRepo "github.com/sagata1999/auth/internal/repository/user/model"
)

const (
	tableName = "\"user\""

	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	roleColumn            = "role"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, create_model *model.CreateUser) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn, passwordConfirmColumn).
		Values(create_model.Name, create_model.Email, create_model.Role, create_model.Password, create_model.PasswordConfirm).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	log.Printf("[repo] create: %s %#v", query, args)

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
