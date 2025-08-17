package infrastructure

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenGenerator interface {
	Generate(identity string, ttl time.Duration) (string, time.Time, error)
}

type JWTTokenGenerator struct {
	secretKey string
}

func NewJWTTokenGenerator(secret string) TokenGenerator {
	return &JWTTokenGenerator{
		secretKey: secret,
	}
}

func (g *JWTTokenGenerator) Generate(identity string, ttl time.Duration) (string, time.Time, error) {
	expire := time.Now().Add(ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": identity,
		"exp": expire.Unix(),
	})

	signedToken, err := token.SignedString([]byte(g.secretKey))
	return signedToken, expire, err
}
