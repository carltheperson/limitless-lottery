package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Get(envName string) string {
	if os.Getenv(envName) != "" {
		return os.Getenv(envName)
	}

	for k, v := range defaults {
		if k == envName {
			return v
		}
	}

	log.Fatalf("The env '%s' was not set and did not have a default", envName)
	return ""
}

var defaults map[string]string = map[string]string{
	"MONGO_URL":    "mongodb://localhost:27017",
	"SERVE_STATIC": "FALSE",
	"SERVE_DIR":    "./ui/webout",
}
