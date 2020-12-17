package auth

import (
	"errors"

	"github.com/carlriis/Limitless-Lottery/db"
	log "github.com/sirupsen/logrus"
)

var ErrLoginAttemptUnsuccessful = errors.New("Login attempt unsuccessful")

func SignIn(username string, password string) (db.SessionIdentity, error) {
	user, err := db.GetUser(username)

	if err != nil {
		return db.SessionIdentity{}, ErrLoginAttemptUnsuccessful
	}

	success := IsPasswordCorrect(password, user.PasswordHash)
	if success != true {
		return db.SessionIdentity{}, ErrLoginAttemptUnsuccessful
	}

	sessionToken := GenerateSessionToken()

	sessionIdentity, err := db.CreateNewSession(username, sessionToken)
	if err != nil {
		log.Errorf("Could not create a session in the db: %v", err)
		return db.SessionIdentity{}, ErrLoginAttemptUnsuccessful
	}

	return sessionIdentity, nil
}
