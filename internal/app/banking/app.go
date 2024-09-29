package banking

import (
	"fmt"
	"net/http"
	"os"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/handler"
	"github.com/quartzeast/go-simple-banking/internal/middleware"
	"github.com/quartzeast/go-simple-banking/internal/pkg/db"
	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
	"github.com/quartzeast/go-simple-banking/internal/service"
	"github.com/quartzeast/go-simple-banking/internal/util"
)

func Start() {
	logger := log.New(log.LevelDebug)
	util.SanityCheck(logger)

	mux := http.NewServeMux()

	// wiring
	dbClient := db.GetDBClient()
	ch := handler.NewCustomerHandler(logger, service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient)))
	ah := handler.NewAccountHandler(logger, service.NewAccountService(domain.NewAccountRepositoryDB(dbClient)))
	am := middleware.NewAuthMiddleware(domain.NewAuthRepository())

	mux.HandleFunc("GET /customers", am.AuthorizationHandler("GetAllCustomers", ch.GetAllCustomer))
	mux.HandleFunc("GET /customers/{id}", am.AuthorizationHandler("GetCustomer", ch.GetCustomer))

	mux.HandleFunc("POST /customers/{id}/account", am.AuthorizationHandler("NewAccount", ah.NewAccount))
	mux.HandleFunc("POST /customers/{customer_id}/account/{account_id}", am.AuthorizationHandler("NewTransaction", ah.MakeTransaction))

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info("Starting HTTP server on port 8090")
	logger.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), mux).Error())
}
