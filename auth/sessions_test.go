package auth_test

import (
	"testing"

	"github.com/carlriis/Limitless-Lottery/auth"
)

func TestSessionTokenGeneration(t *testing.T) {
	if auth.GenerateSessionToken() == auth.GenerateSessionToken() {
		t.Error("auth session tokens the same")
	}
	t.Error(auth.GenerateSessionToken())
}
