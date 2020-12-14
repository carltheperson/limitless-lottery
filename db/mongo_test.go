package db_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/carlriis/Limitless-Lottery/db"
)

func TestDatabaseConnection(t *testing.T) {
	db.Connect()
	defer db.Disconnect()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	got := db.Ping(ctx)

	if got != nil {
		t.Errorf("Could not ping database: %s", got)
	}
}

func TestUserActions(t *testing.T) {
	db.Connect()
	defer db.Disconnect()

	testingUsername := generateTestingUsername()

	err := db.NewUser(testingUsername, "#####")
	if err != nil {
		t.Errorf("Could not create test user '%s' : %s", testingUsername, err)
	}
	user, err := db.GetUser(testingUsername)
	if user.Username != testingUsername || err != nil {
		t.Errorf("Could not retrieve test user '%s' : %s", testingUsername, err)
	}
	err = db.RemoveUser(testingUsername)
	if err != nil {
		t.Errorf("Could not remove test user: '%s'", err)
	}
}

func generateTestingUsername() string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letters := make([]rune, 10)
	for i := range letters {
		letters[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "test_user_" + string(letters)
}
