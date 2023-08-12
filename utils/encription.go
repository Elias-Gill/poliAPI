package utils

import "golang.org/x/crypto/bcrypt"

// GenerateBcryptHash generates a bcrypt hash of the given password string
func EncriptPasw(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CompareBcryptHash compares a bcrypt hash with the given password string
func ComparePasw(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
