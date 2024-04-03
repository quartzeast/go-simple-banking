package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/service"
)

// 5. CustomerHandlers 依赖业务逻辑领域提供的 CustomerService 以完成其功能
// 在结构体字段中使用 service.CustomerService 就相当于定义了一个 contract，
// 要求业务逻辑提供 GetAllCustomers 和 GetCustomer 两项能力
type CustomerHandlers struct {
	service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := c.service.GetAllCustomers(status)

	if e, ok := err.(*errs.AppError); ok {
		writeResponse(w, e.Code, e.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
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
