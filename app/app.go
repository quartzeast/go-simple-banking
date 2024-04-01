package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quartzeast/go-simple-banking/domain"
	"github.com/quartzeast/go-simple-banking/service"
)

func Start() {
	router := mux.NewRouter()
	// 6. wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
