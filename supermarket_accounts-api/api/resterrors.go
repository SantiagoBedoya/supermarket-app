package api

import (
	"net/http"

	"github.com/SantiagoBedoya/supermarket_accounts-api/accounts"
)

type RestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewBadRequestError(msg string) *RestError {
	return &RestError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
	}
}

func ParseError(err error) *RestError {
	switch err {
	case accounts.InvalidUserCredentials:
		return NewBadRequestError(err.Error())
	case accounts.UserAlreadyExistErr:
		return NewBadRequestError(err.Error())
	default:
		return NewInternalServerError("something went wrong, please try again")
	}
}
