package ui

import (
	"net/http"
	"os"

	"github.com/carlriis/Limitless-Lottery/config"
	log "github.com/sirupsen/logrus"
)

// GetStaticHandler gets the handler for the static files that are generated for the frontend. This is for production
func GetStaticHandler() http.Handler {

	_, err := os.Stat(config.Get("SERVE_DIR"))
	if err != nil {
		log.Fatalf("The directory '%s' does not exist, but is needed for Serve", config.Get("SERVE_DIR"))
	}

	fs := http.FileServer(http.Dir(config.Get("SERVE_DIR")))
	return fs
}
