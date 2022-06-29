package products

import "errors"

var (
	InvalidProductDataErr   = errors.New("invalid product data")
	ProductAlreadyExistsErr = errors.New("this code is already in use for other product")
	ProductNotFoundErr      = errors.New("product not found")
)
