package helper

import (

	"golang.org/x/crypto/bcrypt"
) 



func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost) 
	if err != nil {
		return "", err
	} 
	return string(hash), nil
}

func VerifyPassword(passhash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password)) 
	if err != nil {
		return err
	}
	return nil 
}