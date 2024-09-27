package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/quartzeast/go-simple-banking/internal/pkg/apierr"
	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
	"github.com/quartzeast/go-simple-banking/internal/response"
	"github.com/quartzeast/go-simple-banking/internal/service"
)

type CustomerHandler struct {
	logger  *log.Logger
	service service.CustomerService
}

func NewCustomerHandler(logger *log.Logger, service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		logger:  logger,
		service: service,
	}
}

func (h *CustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetAllCustomer()
	if err != nil {
		h.logger.Error("get all customer failed", "error", err.Error())
		response.Error(w, apierr.NewAPIError(apierr.CodeUnknownError, err))
		return
	}

	response.OK(w, http.StatusOK, customers)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	customer, err := h.service.GetCustomer(id)
	if err != nil {
		h.logger.Error("get customer failed", "id", id, "error", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, apierr.NewAPIError(apierr.CodeNotFound, err))
			return
		}

		response.Error(w, apierr.NewAPIError(apierr.CodeUnknownError, err))
		return
	}

	response.OK(w, http.StatusOK, customer)
}
