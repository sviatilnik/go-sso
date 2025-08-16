package user

import (
	"context"
	"time"
)

type Gender string

var Male = Gender("male")
var Female = Gender("female")

type Profile struct {
	ID        string
	UserID    string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Gender    Gender
	Created   time.Time
	Updated   time.Time
}

type ProfileRepository interface {
	FindByUserID(ctx context.Context, ID string) (*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, ID string) error
}
