package auth

import (
	"net/http"

	"github.com/carlriis/Limitless-Lottery/db"
)

func LogOut(r *http.Request) {
	cookie, err := r.Cookie("session_token")

	if err != nil {
		return
	}

	sessionToken := cookie.Value

	db.RevokeSession(sessionToken)
}
