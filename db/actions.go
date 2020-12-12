package db

import (
	"errors"
)

const initialBalance = 10000

var ErrUsernameTaken = errors.New("Username taken")
var ErrUserDoesNotExist = errors.New("User does not exist")

func NewUser(username string) error {
	ctx, cancel := getContext()
	defer cancel()

	if checkIfUsernameTaken(username) {
		return ErrUsernameTaken
	}

	_, err := usersCollection.InsertOne(ctx, User{Username: username, Balance: initialBalance})
	if err != nil {
		return err
	}

	return nil
}

func RemoveUser(username string) error {
	ctx, cancel := getContext()
	defer cancel()

	result, _ := usersCollection.DeleteOne(ctx, User{Username: username})

	if result.DeletedCount == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}

func GetUser(username string) (User, error) {
	ctx, cancel := getContext()
	defer cancel()

	var result User
	usersCollection.FindOne(ctx, User{Username: username}).Decode(&result)

	if result.Username == "" {
		return User{}, ErrUserDoesNotExist
	}

	return result, nil
}

func ChangeUserBalance(username string, change int) error {
	user, err := GetUser(username)
	if err != nil {
		return err
	}

	user.Balance += change

	ctx, cancel := getContext()
	defer cancel()
	_, err = usersCollection.UpdateOne(ctx, User{Username: username}, user)

	return err
}

func checkIfUsernameTaken(username string) bool {
	ctx, cancel := getContext()
	defer cancel()

	count, _ := usersCollection.CountDocuments(ctx, User{Username: username})

	if count == 0 {
		return false
	}
	return true
}
