package auth

import "golang.org/x/crypto/bcrypt"

type sessionToken string

const pepper string = "Gnfguw6QKJA2TyzRFEYQ"

func GenerateHashedPassword(password string) string {
	pepperedPassword := []byte(password + pepper)

	hashedPassword, _ := bcrypt.GenerateFromPassword(pepperedPassword, 10)

	return string(hashedPassword)
}

func IsPasswordCorrect(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+pepper))
	if err == nil {
		return true
	}
	return false
}
