package apierr

import (
	"encoding/json"
	"fmt"
)

func NewAPIError(coder APICoder, cause ...error) error {
	var c error
	if len(cause) > 0 {
		c = cause[0]
	}

	return &apiError{
		coder: coder,
		cause: c,
	}
}

type apiError struct {
	coder APICoder
	cause error
}

func (a *apiError) Error() string {
	return fmt.Sprintf("[%d] - %s", a.coder.Code(), a.coder.Message())
}

func (a *apiError) Coder() APICoder {
	return a.coder
}

func (a *apiError) Unwrap() error {
	return a.cause
}

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

func (a *apiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&errorMessage{
		Code:    a.coder.Code(),
		Message: a.coder.Message(),
	})
}

func (a *apiError) UnmarshalJSON(data []byte) error {
	e := &errorMessage{}
	if err := json.Unmarshal(data, e); err != nil {
		return err
	}
	a.coder = NewAPICode(e.Code, e.Message)
	return nil
}
