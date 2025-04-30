package domain

type PasswordHash interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}
