package user

import (
	"context"
	"time"
)

type User struct {
	ID       string
	Active   bool
	Login    string
	Password string
	Profile  *Profile
	Roles    []*Role
	Created  time.Time
	Updated  time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByLogin(ctx context.Context, login string) (*User, error)
}
