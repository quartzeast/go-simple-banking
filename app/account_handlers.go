package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quartzeast/go-simple-banking/dto"
	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (a AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerID
	account, err := a.service.NewAccount(request)
	if e, ok := err.(*errs.AppError); ok {
		writeResponse(w, e.Code, e)
		return
	}
	writeResponse(w, http.StatusCreated, account)
}
