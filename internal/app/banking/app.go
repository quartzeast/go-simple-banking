package banking

import (
	"fmt"
	"net/http"
	"os"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/handler"
	"github.com/quartzeast/go-simple-banking/internal/pkg/db"
	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
	"github.com/quartzeast/go-simple-banking/internal/service"
)

func sanityCheck(logger *log.Logger) {
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

func Start() {
	logger := log.New(log.LevelDebug)
	sanityCheck(logger)

	mux := http.NewServeMux()

	// wiring
	dbClient := db.GetDBClient()
	ch := handler.NewCustomerHandler(logger, service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient)))
	mux.HandleFunc("GET /customers", ch.GetAllCustomer)
	mux.HandleFunc("GET /customers/{id}", ch.GetCustomer)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info("Starting HTTP server on port 8090")
	logger.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), mux).Error())
}
