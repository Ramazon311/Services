package etc

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash ...
func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}

func CompareHashAndPassword(hashedPassword, password []byte) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
}