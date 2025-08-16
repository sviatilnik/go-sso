package user

import (
	"context"
	"time"
)

type UserRole struct {
	UserID  string
	RoleID  string
	Created time.Time
	Start   time.Time
	End     time.Time
}

type UserRoleRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*UserRole, error)
	AddRoleToUser(ctx context.Context, userID, roleID string) error
	RemoveRoleFromUser(ctx context.Context, userID, roleID string) error
	UserHasRole(ctx context.Context, userID, roleID string) (bool, error)
}
