package infrastructure

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) bool
}

type DefaultHasher struct{}

func NewDefaultHasher() Hasher {
	return &DefaultHasher{}
}

func (h DefaultHasher) Hash(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(fromPassword), nil
}

func (h DefaultHasher) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
