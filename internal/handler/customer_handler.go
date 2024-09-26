package handler

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/quartzeast/go-simple-banking/internal/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (ch *CustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
