package apierr

import "errors"

// 1-3: HTTP status code
// 4-6: Module code
// 7-9: Error code
var (
	CodeBadRequest   = NewAPICode(40000000, "invalid request")
	CodeUnauthorized = NewAPICode(40100000, "unauthorized request")
	CodeForbidden    = NewAPICode(40300000, "forbidden request")
	CodeNotFound     = NewAPICode(40400000, "resource not found")
	CodeUnknownError = NewAPICode(50000000, "unknown error")
)

type APICoder interface {
	Code() int
	Message() string
	HTTPStatusCode() int
}

func NewAPICode(code int, message string) APICoder {
	return &apiCode{
		code:    code,
		message: message,
	}
}

type apiCode struct {
	code    int
	message string
}

func (a *apiCode) Code() int {
	return a.code
}

func (a *apiCode) Message() string {
	return a.message
}

func (a *apiCode) HTTPStatusCode() int {
	v := a.Code()
	for v >= 1000 {
		v /= 10
	}
	return v
}

func ParseCoder(err error) APICoder {
	for {
		if e, ok := err.(interface {
			Coder() APICoder
		}); ok {
			return e.Coder()
		}
		if errors.Unwrap(err) == nil {
			return CodeUnknownError
		}
		err = errors.Unwrap(err)
	}
}

func IsCode(err error, coder APICoder) bool {
	if err == nil {
		return false
	}

	for {
		if e, ok := err.(interface {
			Coder() APICoder
		}); ok {
			if e.Coder().Code() == coder.Code() {
				return true
			}
		}

		if errors.Unwrap(err) == nil {
			return false
		}

		err = errors.Unwrap(err)
	}
}
