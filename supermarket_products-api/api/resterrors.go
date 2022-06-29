package api

import (
	"net/http"

	"github.com/SantiagoBedoya/supermarket_products-api/products"
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

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
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
	case products.ProductAlreadyExistsErr:
		return NewBadRequestError(err.Error())
	case products.ProductNotFoundErr:
		return NewNotFoundError(err.Error())
	case products.InvalidProductDataErr:
		return NewBadRequestError(err.Error())
	default:
		return NewInternalServerError("something went wrong, please try again")
	}
}
