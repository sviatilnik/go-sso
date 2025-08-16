package user

import (
	"context"
	"time"
)

type Role struct {
	ID          string
	Name        string
	Description string
	Created     time.Time
	Updated     time.Time
}

type RoleRepository interface {
	Create(ctx context.Context, role *Role) (*Role, error)
	Update(ctx context.Context, role *Role) (*Role, error)
	FindByID(ctx context.Context, id string) (*Role, error)
	FindByName(ctx context.Context, name string) (*Role, error)
	FindAll(ctx context.Context) ([]*Role, error)
}
