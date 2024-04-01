package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipcode" xml:"zipcode"`
}

// 5. CustomerHandlers 依赖 Service Port
type CustomerHandlers struct {
	service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := c.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (c *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := c.service.GetCustomer(id)
	if e, ok := err.(*errs.AppError); ok {
		writeResponse(w, e.Code, e)
		return
	}

	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
