package user

import "time"

type TokenPair struct {
	AccessToken  Token
	RefreshToken Token
}

type Token struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
}

func (t *Token) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}
