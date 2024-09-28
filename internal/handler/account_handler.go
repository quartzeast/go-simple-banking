package handler

import (
	"encoding/json"
	"net/http"

	"github.com/quartzeast/go-simple-banking/internal/dto"
	"github.com/quartzeast/go-simple-banking/internal/pkg/apierr"
	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
	"github.com/quartzeast/go-simple-banking/internal/response"
	"github.com/quartzeast/go-simple-banking/internal/service"
)

type AccountHandler struct {
	logger  *log.Logger
	service service.AccountService
}

func NewAccountHandler(logger *log.Logger, service service.AccountService) *AccountHandler {
	return &AccountHandler{
		logger:  logger,
		service: service,
	}
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("id")
	var request dto.NewAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Error(err.Error())
		response.Error(w, apierr.NewAPIError(apierr.CodeBadRequest, err))
		return
	}

	request.CustomerId = customerID
	account, err := h.service.NewAccount(request)
	if err != nil {
		h.logger.Error(err.Error())
		response.Error(w, apierr.NewAPIError(apierr.CodeBadRequest, err))
		return
	}

	response.OK(w, http.StatusCreated, account)
}
