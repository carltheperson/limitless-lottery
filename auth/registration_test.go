package auth_test

import (
	"testing"

	"github.com/carlriis/Limitless-Lottery/auth"
)

func TestPasswordComparison(t *testing.T) {
	const password string = "supe3r_g00d_p4ssw0rd_123"

	hash := auth.GenerateHashedPassword(password)

	isPasswordCorrect := auth.IsPasswordCorrect(password, hash)
	if !isPasswordCorrect {
		t.Error("auth thought that correct password was not correct")
	}

	isPasswordCorrect = auth.IsPasswordCorrect(password+"_not_correct", hash)
	if isPasswordCorrect {
		t.Error("auth thought that incorrect password was correct")
	}

}
