package postgres

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/sviatilnik/sso/internal/sso/domain/user"
)

type UserRepository struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (u *UserRepository) Create(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Update(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	query, _, err := u.builder.Select("id", "login", "password").
		From("users").
		Where("id = ?").
		ToSql()

	if err != nil {
		return nil, err
	}

	_user := new(user.User)

	err = u.db.QueryRowContext(ctx, query, id).Scan(_user.ID, _user.Login, _user.Password)
	if err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *UserRepository) FindByLogin(ctx context.Context, login string) (*user.User, error) {
	query, _, err := u.builder.Select("id", "login", "password").
		From("users").
		Where("login = ?").
		ToSql()

	if err != nil {
		return nil, err
	}

	_user := new(user.User)

	err = u.db.QueryRowContext(ctx, query, login).Scan(_user.ID, _user.Login, _user.Password)
	if err != nil {
		return nil, err
	}

	return _user, nil
}
