package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/carlriis/Limitless-Lottery/db"
)

var (
	ErrRetrievingSessionToken = errors.New("Could not retrieve a session cookie")
	ErrInvalidSessionToken    = errors.New("Session token is invalid")
)

func Authenticate(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")

	if err != nil {
		return "", ErrRetrievingSessionToken
	}

	sessionToken := cookie.Value

	sessionIdentity, err := db.FindSessionIdentityFromSessionToken(sessionToken)
	if err == db.ErrThereIsNoSessionIdentityForThatSessionToken {
		return "", ErrInvalidSessionToken
	}

	if time.Now().Unix() > sessionIdentity.ExpirationDate {
		db.RevokeSession(sessionIdentity.Username)
		return "", ErrInvalidSessionToken
	}

	return sessionIdentity.Username, nil
}
