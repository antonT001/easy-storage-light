package httplib

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	RestError struct {
		Code       string   `json:"code" validate:"required"`
		Message    string   `json:"message" validate:"required"`
		Args       []string `json:"args,omitempty"`
		Err        error    `json:"-"`
		HTTPStatus int      `json:"-"`
	}
	ErrorRegistry map[string]RestError
)

const (
	GeneralError              = "GeneralError"
	GeneralFileMgrError       = "GeneralFileMgrError"
	InvalidObject             = "InvalidObject"
	InvalidParam              = "InvalidParam"
	ErrorParsingServiceResult = "ErrorParsingServiceResult"
	ErrorSendingRequest       = "ErrorSendingRequest"
	PayloadError              = "PayloadError"
	RequestTimeout            = "RequestTimeout"
)

var baseErrorRegistry = ErrorRegistry{
	GeneralError: RestError{
		Message:    "unexpected error happen, please contact the support",
		HTTPStatus: http.StatusInternalServerError,
	},
	GeneralFileMgrError: RestError{
		Message:    "operation failed for %1",
		HTTPStatus: http.StatusInternalServerError,
	},
	ErrorParsingServiceResult: RestError{
		Message:    "error while getting results",
		HTTPStatus: http.StatusBadGateway,
	},
	ErrorSendingRequest: RestError{
		Message:    "error sending request",
		HTTPStatus: http.StatusInternalServerError,
	},
	InvalidObject: RestError{
		Message:    "provided object is invalid: %1",
		HTTPStatus: http.StatusUnprocessableEntity,
	},
	InvalidParam: RestError{
		Message:    "provided parameter %1 is invalid",
		HTTPStatus: http.StatusBadRequest,
	},
	PayloadError: RestError{
		Message:    "invalid payload: %1",
		HTTPStatus: http.StatusBadRequest,
	},
	RequestTimeout: RestError{
		Message:    "request timeout: %1",
		HTTPStatus: http.StatusRequestTimeout,
	},
}

func (e *RestError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e *RestError) Is(target error) bool {
	var restErr *RestError
	if !errors.As(target, &restErr) {
		return false
	}
	return e.Code != "" && e.Code == restErr.Code
}

func (e *RestError) Unwrap() error {
	return e.Err
}

func NewError(e error, code string, args ...interface{}) *RestError {
	strArgs := make([]string, 0, len(args))
	for i := range args {
		strArgs = append(strArgs, fmt.Sprint(args[i]))
	}

	restErr, exist := baseErrorRegistry[code]
	if exist {
		restErr.Code = code
	} else {
		restErr = baseErrorRegistry[GeneralError]
		restErr.Code = GeneralError
	}
	restErr.Args = strArgs
	restErr.Err = e
	return &restErr
}

func AsRestErr(err error) *RestError {
	if err == nil {
		return nil
	}
	var restErr *RestError
	if !errors.As(err, &restErr) {
		restErr = NewError(err, GeneralError)
	}
	return restErr
}
