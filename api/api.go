package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Serve starts the API
func Serve(addr string) {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(jsonMiddleware)

	r.HandleFunc("/checkticketamount", checkTicketAmount).Methods("PUT")
	r.HandleFunc("/checkticketuntilwin", checkTicketUntilWin).Methods("PUT")
	r.HandleFunc("/signin", signIn).Methods("POST")
	r.HandleFunc("/signup", signUp).Methods("POST")

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	log.Info("Starting server on " + addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
