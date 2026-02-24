package infrastructure

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) bool
}

type BcryptHasher struct{}

func NewBcryptHasher() Hasher {
	return &BcryptHasher{}
}

func (h BcryptHasher) Hash(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(fromPassword), nil
}

func (h BcryptHasher) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
