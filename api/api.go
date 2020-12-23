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

	rServe.Use(corsMiddleware)

	if config.Get("SERVE_STATIC") == "TRUE" {
		fn := ui.GetStaticHandler()
		rServe.PathPrefix("/").Handler(http.StripPrefix("/", fn))
	}

	r := rServe.PathPrefix("/api").Subrouter()

	r.Use(jsonMiddleware)
	r.Use(corsMiddleware)

	r.Methods("OPTIONS").HandlerFunc(corsOptions)

	r.HandleFunc("/checkticketamount", checkTicketAmount).Methods("PUT")
	r.HandleFunc("/checkticketuntilwin", checkTicketUntilWin).Methods("PUT")
	r.HandleFunc("/signin", signIn).Methods("POST")
	r.HandleFunc("/signup", signUp).Methods("POST")
	r.HandleFunc("/session-username", retreiveUsername).Methods("GET")

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.Get("ALLOW_ORIGIN_URL"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Method", "*")
		next.ServeHTTP(w, r)
	})
}

func corsOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", config.Get("ALLOW_ORIGIN_URL"))
	w.Header().Set("Access-Control-Request-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}
