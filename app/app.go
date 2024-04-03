package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/quartzeast/go-simple-banking/domain"
	"github.com/quartzeast/go-simple-banking/logger"
	"github.com/quartzeast/go-simple-banking/service"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()
	// 6. wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// REST Handlers，负责驱动业务逻辑，依赖 service
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepository())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
