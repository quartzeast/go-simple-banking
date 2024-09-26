package banking

import (
	"log"
	"net/http"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/handler"
	"github.com/quartzeast/go-simple-banking/internal/service"
)

func Start() {
	mux := http.NewServeMux()

	// wiring
	ch := handler.NewCustomerHandler(service.NewCustomerService(domain.NewCustomerRepositoryStub()))
	mux.HandleFunc("/customers", ch.GetAllCustomer)

	log.Println("Starting HTTP server on port 8090")
	log.Fatalln(http.ListenAndServe(":8090", mux))
}
