package response

import (
	"encoding/json"
	"net/http"

	"github.com/quartzeast/go-simple-banking/internal/pkg/apierr"
)

func writeResponse(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func OK(w http.ResponseWriter, status int, result any) {
	if result == nil {
		writeResponse(w, http.StatusNoContent, nil)
		return
	}
	writeResponse(w, status, result)
}

func Error(w http.ResponseWriter, err error) {
	e := apierr.ParseCoder(err)
	statusCode := e.HTTPStatusCode()
	writeResponse(w, statusCode, err)
}
