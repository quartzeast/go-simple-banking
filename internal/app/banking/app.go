package banking

import (
	"log"
	"net/http"
	"os"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/handler"
	"github.com/quartzeast/go-simple-banking/internal/pkg/db"
	"github.com/quartzeast/go-simple-banking/internal/service"
)

func sanityCheck() {
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
			log.Fatalf("Environment variable %s not defined. Terminating application...", env)
		}
	}
}

func Start() {
	sanityCheck()

	mux := http.NewServeMux()

	// wiring
	dbClient := db.GetDBClient()
	ch := handler.NewCustomerHandler(service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient)))
	mux.HandleFunc("/customers", ch.GetAllCustomer)

	log.Println("Starting HTTP server on port 8090")
	log.Fatalln(http.ListenAndServe(":8090", mux))
}
