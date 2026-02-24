package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenGenerator interface {
	Generate(identity string, ttl time.Duration) (string, time.Time, error)
}

type JWTTokenGenerator struct {
	secretKey string
	issuer    string
	kid       string
}

func NewJWTTokenGenerator(secret, kid, issuer string) TokenGenerator {
	return &JWTTokenGenerator{
		secretKey: secret,
		issuer:    issuer,
		kid:       kid,
	}
}

func (g *JWTTokenGenerator) Generate(identity string, ttl time.Duration) (string, time.Time, error) {
	expire := time.Now().Add(ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"sub": identity,
		"exp": expire.Unix(),
	})

	token.Header["kid"] = g.kid

	signedToken, err := token.SignedString([]byte(g.secretKey))
	return signedToken, expire, err
}
