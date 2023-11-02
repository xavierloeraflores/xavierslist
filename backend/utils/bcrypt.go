package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash
}

func ComparePassword(password string, hash []byte) bool {
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(hash, bytePassword)
	if err != nil {
		return false
	}
	return true
}
