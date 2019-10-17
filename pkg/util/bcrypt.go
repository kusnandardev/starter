package util

import "golang.org/x/crypto/bcrypt"

// Hash :
func Hash(text string) (string, error) {
	pwd := []byte(text)

	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return text, err
	}
	return string(hashedPwd), nil
}

// Compare :
func Compare(hashed string, text string) (bool, error) {
	hashedPwd := []byte(hashed)
	pwd := []byte(text)

	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)
	if err != nil {
		return false, err
	}
	return true, nil
}
