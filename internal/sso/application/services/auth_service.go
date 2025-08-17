package services

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/sviatilnik/sso/internal/sso/application"
	"github.com/sviatilnik/sso/internal/sso/domain/user"
	"github.com/sviatilnik/sso/internal/sso/infrastructure"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest *application.LoginRequest) (*application.LoginResponse, error)
}

type AuthServiceImpl struct {
	userRepo  user.UserRepository
	hasher    infrastructure.Hasher
	generator infrastructure.TokenGenerator
}

func NewAuthService(userRepo user.UserRepository, hasher infrastructure.Hasher, generator infrastructure.TokenGenerator) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		hasher:    hasher,
		generator: generator,
	}
}

func (i *AuthServiceImpl) Login(ctx context.Context, loginRequest *application.LoginRequest) (*application.LoginResponse, error) {
	u, err := i.userRepo.FindByLogin(ctx, loginRequest.Login)
	if err != nil {
		return nil, err
	}

	if !u.Active {
		return nil, errors.New("user is not active")
	}

	isCorrectPwd := i.hasher.Compare(u.Password, loginRequest.Password)
	if !isCorrectPwd {
		return nil, err
	}

	return i.generateTokens(u.ID)
}

func (i *AuthServiceImpl) generateTokens(userID string) (*application.LoginResponse, error) {
	access, exp, err := i.generator.Generate(userID, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	accessToken := application.Token{
		Token:  access,
		Expire: exp.Unix(),
	}

	h := sha256.New()
	h.Write([]byte(access))

	refreshToken := application.Token{
		Token:  base64.RawURLEncoding.EncodeToString(h.Sum(nil)),
		Expire: time.Now().Add(24 * time.Hour).Unix(),
	}

	return &application.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
