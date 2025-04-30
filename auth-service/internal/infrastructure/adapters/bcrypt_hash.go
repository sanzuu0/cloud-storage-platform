package adapters

import "golang.org/x/crypto/bcrypt"

type BcryptHash struct{}

func NewBcryptHash() *BcryptHash {
	return &BcryptHash{}
}

func (BcryptHash) Hash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func (BcryptHash) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
