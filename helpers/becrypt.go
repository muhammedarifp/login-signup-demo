package helpers

import "golang.org/x/crypto/bcrypt"

func ToHash(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHash(hash string, inp string) bool {
	hashbyte := []byte(hash)
	inpbyte := []byte(inp)

	err := bcrypt.CompareHashAndPassword(hashbyte, inpbyte)
	if err == nil {
		return true
	}
	return false
}
