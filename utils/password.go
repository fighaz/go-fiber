package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func CheckPassword(password string, hashPass string) bool {
	valid := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

	return valid == nil
}
