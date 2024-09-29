package util

import (
	"os"

	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
)

func SanityCheck(logger *log.Logger) {
	envs := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, env := range envs {
		if os.Getenv(env) == "" {
			logger.Fatal("Environment variable not defined. Terminating application...", "env", env)
		}
	}
}
