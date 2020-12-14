package db

import (
	"errors"
	"time"
)

const daysTillExpiration = 30

var ErrThereIsNoSessionTokenForThatUser = errors.New("There is no session token for that user")

func CreateNewSession(username string, sessionToken string) error {
	ctx, cancel := getContext()
	defer cancel()

	_, err := sessionIdentitiesCollection.InsertOne(ctx, SessionIdentity{
		Username:       username,
		SessionToken:   sessionToken,
		ExpirationDate: time.Now().AddDate(0, 0, daysTillExpiration).Unix(),
	})

	return err
}

func RevokeSession(username string) error {
	ctx, cancel := getContext()
	defer cancel()

	_, err := sessionIdentitiesCollection.DeleteOne(ctx, SessionIdentity{Username: username})

	return err
}

func RetrieveSessionToken(username string) (string, error) {
	ctx, cancel := getContext()
	defer cancel()

	var result SessionIdentity
	sessionIdentitiesCollection.FindOne(ctx, SessionIdentity{Username: username}).Decode(&result)

	if result.SessionToken == "" {
		return "", ErrThereIsNoSessionTokenForThatUser
	}

	return result.SessionToken, nil
}
