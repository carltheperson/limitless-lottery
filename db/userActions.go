package db

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const initialBalance = 1000

var ErrUsernameTaken = errors.New("Username taken")
var ErrUserDoesNotExist = errors.New("User does not exist")

func NewUser(username string, passwordHash string) error {
	ctx, cancel := getContext()
	defer cancel()

	_, err := usersCollection.InsertOne(ctx, User{
		Username:     username,
		PasswordHash: passwordHash,
		Balance:      initialBalance,
	})

	return err
}

func RemoveUser(username string) error {
	ctx, cancel := getContext()
	defer cancel()

	result, err := usersCollection.DeleteOne(ctx, User{Username: username})
	if err != nil {
		log.Error(err)
	}

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

func ChangeUserBalance(username string, change int) (int, error) {
	user, err := GetUser(username)
	if err != nil {
		return 0, err
	}

	user.Balance += change

	ctx, cancel := getContext()
	defer cancel()

	update := bson.M{
		"$set": user,
	}

	_, err = usersCollection.UpdateOne(ctx, User{Username: username}, update)
	if err != nil {
		log.Error(err)
	}

	return user.Balance, err
}

func CheckIfUsernameTaken(username string) bool {
	ctx, cancel := getContext()
	defer cancel()

	count, err := usersCollection.CountDocuments(ctx, User{Username: username})
	if err != nil {
		log.Error(err)
	}

	if count == 0 {
		return false
	}
	return true
}
