package db

import (
	"errors"
	"time"
)

const daysTillExpiration = 30

var (
	ErrThereIsNoSessionTokenForThatUser            = errors.New("There is no session token for that user")
	ErrThereIsNoSessionIdentityForThatSessionToken = errors.New("There is no session identity for that session token")
)

func CreateNewSession(username string, sessionToken string) (SessionIdentity, error) {
	ctx, cancel := getContext()
	defer cancel()

	sessionIdentity := SessionIdentity{
		Username:       username,
		SessionToken:   sessionToken,
		ExpirationDate: time.Now().AddDate(0, 0, daysTillExpiration).Unix(),
	}

	_, err := sessionIdentitiesCollection.InsertOne(ctx, sessionIdentity)

	return sessionIdentity, err
}

// RevokeSession removes the SessionIdentity  (in the db) that contains the provided session token
func RevokeSession(sessionToken string) error {
	ctx, cancel := getContext()
	defer cancel()

	_, err := sessionIdentitiesCollection.DeleteOne(ctx, SessionIdentity{SessionToken: sessionToken})

	return err
}

func FindSessionIdentityFromSessionToken(sessionToken string) (SessionIdentity, error) {
	ctx, cancel := getContext()
	defer cancel()

	var result SessionIdentity
	sessionIdentitiesCollection.FindOne(ctx, SessionIdentity{SessionToken: sessionToken}).Decode(&result)

	if result.Username == "" {
		return SessionIdentity{}, ErrThereIsNoSessionIdentityForThatSessionToken
	}

	return result, nil
}
