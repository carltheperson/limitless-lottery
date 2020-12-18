package auth

import (
	"github.com/carlriis/Limitless-Lottery/db"
	log "github.com/sirupsen/logrus"
)

// SignUp stores a user in the db with the password in hashed form
func SignUp(username string, password string) error {
	passwordHash := GenerateHashedPassword(password)

	err := db.NewUser(username, passwordHash)
	if err != nil {
		log.Errorf("Could not create new user: %v", err)
		return err
	}

	return nil
}
