package api

import (
	"net/http"
	"time"

	"github.com/carlriis/Limitless-Lottery/config"
	"github.com/carlriis/Limitless-Lottery/ui"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Serve starts the API
func Serve(addr string) {
	rServe := mux.NewRouter()

	if config.Get("SERVE_STATIC") == "TRUE" {
		fn := ui.GetStaticHandler()
		rServe.Handle("/", fn)
	}

	r := rServe.PathPrefix("/api").Subrouter()

	r.Use(jsonMiddleware)

	r.HandleFunc("/checkticketamount", checkTicketAmount).Methods("PUT")
	r.HandleFunc("/checkticketuntilwin", checkTicketUntilWin).Methods("PUT")
	r.HandleFunc("/signin", signIn).Methods("POST")
	r.HandleFunc("/signup", signUp).Methods("POST")

	server := &http.Server{
		Addr:         addr,
		Handler:      rServe,
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
