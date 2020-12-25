package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/carlriis/Limitless-Lottery/api/validation"
	"github.com/carlriis/Limitless-Lottery/auth"
	"github.com/carlriis/Limitless-Lottery/db"
	"github.com/carlriis/Limitless-Lottery/tickets"
)

var (
	errNoTicketMatch        = validation.ErrorMessage{Case: "match", Field: "ID", Message: "There was no match for that ticket id"}
	errUserDoesNotExist     = validation.ErrorMessage{Case: "exist", Field: "Username", Message: "No user was found with that username"}
	errCouldNotAuthenticate = validation.ErrorMessage{Case: "auth", Field: "session_token", Message: "Could not authenticate you"}
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()

	var input struct {
		ID     string `json:"ticketid" validate:"required"`
		Amount int    `json:"amount" validate:"required,numeric,min=0,max=1000000000"`
	}

	validation.UnmarshalJSONAndAddErrors(&input, r.Body, &ea)
	if ea.HasErrors() {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	username, err := auth.Authenticate(r)
	if err != nil {
		ea.Add(errCouldNotAuthenticate)
		ea.Flush(w, http.StatusForbidden)
		return
	}

	ct, err := tickets.CheckAmount(input.ID, input.Amount)

	if err == tickets.ErrIDWithNoMatch {
		ea.Add(errNoTicketMatch)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(username, ct.AmountWonTotal-ct.AmountDeducted)

	if err == db.ErrUserDoesNotExist {
		ea.Add(errUserDoesNotExist)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Ct      tickets.CheckedTicketAmount
		Balance int
	}{
		Ct:      ct,
		Balance: updatedBalance,
	})
}

func checkTicketUntilWin(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()

	var input struct {
		ID string `json:"ticketid" validate:"required"`
	}

	validation.UnmarshalJSONAndAddErrors(&input, r.Body, &ea)
	if ea.HasErrors() {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	username, err := auth.Authenticate(r)
	if err != nil {
		ea.Add(errCouldNotAuthenticate)
		ea.Flush(w, http.StatusForbidden)
		return
	}

	ct, err := tickets.CheckUntilWin(input.ID)

	if err == tickets.ErrIDWithNoMatch {
		ea.Add(errNoTicketMatch)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(username, ct.Profit)

	if err == db.ErrUserDoesNotExist {
		ea.Add(errUserDoesNotExist)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Ct      tickets.CheckedTicketUntilWin
		Balance int
	}{
		Ct:      ct,
		Balance: updatedBalance,
	})
}

func signIn(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()

	var input struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	validation.UnmarshalJSONAndAddErrors(&input, r.Body, &ea)
	if ea.HasErrors() {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	sessionIdentity, err := auth.SignIn(input.Username, input.Password)

	if err != nil {
		ea.Add(validation.ErrorMessage{Case: "login", Field: "Username||Password", Message: "Could not log you in"})
		ea.Flush(w, http.StatusForbidden)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionIdentity.SessionToken,
		Expires:  time.Unix(sessionIdentity.ExpirationDate, 0),
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()

	var input struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	validation.UnmarshalJSONAndAddErrors(&input, r.Body, &ea)
	if ea.HasErrors() {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	taken := db.CheckIfUsernameTaken(input.Username)
	if taken {
		ea.Add(validation.ErrorMessage{Case: "taken", Field: "username", Message: "The username '" + input.Username + "' is taken"})
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	err := auth.SignUp(input.Username, input.Password)

	if err != nil {
		ea.Add(validation.ErrorMessage{Case: "sign up", Field: "", Message: "Could not sign you up"})
		ea.Flush(w, http.StatusInternalServerError)
	}

	// Creating session
	sessionIdentity, err := auth.SignIn(input.Username, input.Password)

	if err != nil {
		ea.Add(validation.ErrorMessage{Case: "login", Field: "Username||Password", Message: "Could not log you in"})
		ea.Flush(w, http.StatusForbidden)
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionIdentity.SessionToken,
		Expires:  time.Unix(sessionIdentity.ExpirationDate, 0),
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
}

func retreiveUsername(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()
	username, err := auth.Authenticate(r)
	if err != nil {
		ea.Add(errCouldNotAuthenticate)
		ea.Flush(w, http.StatusForbidden)
		return
	}
	w.Write([]byte(username))
}

func retrieveBalance(w http.ResponseWriter, r *http.Request) {
	username, err := auth.Authenticate(r)
	if err != nil {
		return
	}
	user, _ := db.GetUser(username)

	w.Write([]byte(strconv.Itoa(user.Balance)))
}

func retrieveOdds(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Odds []string
	}{
		Odds: tickets.ExportOdds(),
	})
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	auth.LogOut(r)

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
}
