package auth

import "golang.org/x/crypto/bcrypt"

type sessionToken string

const pepper string = "Gnfguw6QKJA2TyzRFEYQ" // Changing this will break the password checking for existing users

// GenerateHashedPassword creates a new hashed password
func GenerateHashedPassword(password string) string {
	pepperedPassword := []byte(password + pepper)

	hashedPassword, _ := bcrypt.GenerateFromPassword(pepperedPassword, 10)

	return string(hashedPassword)
}

// IsPasswordCorrect compares a password with a hashedPassword
func IsPasswordCorrect(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+pepper))
	if err == nil {
		return true
	}
	return false
}
