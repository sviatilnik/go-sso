package postgres

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/sviatilnik/sso/internal/sso/domain/user"
)

type RoleRepository struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func (r *RoleRepository) Create(ctx context.Context, role *user.Role) (*user.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) Update(ctx context.Context, role *user.Role) (*user.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) FindByID(ctx context.Context, id string) (*user.Role, error) {
	query, _, err := r.builder.Select("id", "name", "description", "created", "updated").From("roles").Where("id = ?").ToSql()

	if err != nil {
		return nil, err
	}

	role := new(user.Role)

	err = r.db.QueryRowContext(ctx, query, id).Scan(role.ID, role.Name, role.Description, role.Created, role.Updated)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*user.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) GetAll(ctx context.Context) ([]*user.Role, error) {
	panic("implement me")
}
