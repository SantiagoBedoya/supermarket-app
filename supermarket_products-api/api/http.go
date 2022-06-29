package api

import (
	"net/http"

	"github.com/SantiagoBedoya/supermarket_products-api/products"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByCode(c *gin.Context)
	UpdateByCode(c *gin.Context)
	DeleteByCode(c *gin.Context)
}

type handler struct {
	service products.Service
}

func NewHandler(service products.Service) Handler {
	return &handler{service}
}

func (h *handler) Create(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, RestError{
			Message:    products.InvalidProductDataErr.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	newProduct, err := h.service.Create(&product)
	if err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}
func (h *handler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(http.StatusOK, products)
}
func (h *handler) GetByCode(c *gin.Context) {
	code := c.Param("code")
	product, err := h.service.GetByCode(code)
	if err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(http.StatusOK, product)
}
func (h *handler) UpdateByCode(c *gin.Context) {
	code := c.Param("code")
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, RestError{
			Message:    products.InvalidProductDataErr.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	if _, err := h.service.UpdateByCode(code, &product); err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.Status(http.StatusNoContent)
}
func (h *handler) DeleteByCode(c *gin.Context) {
	code := c.Param("code")
	if err := h.service.DeleteByCode(code); err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.Status(http.StatusNoContent)
}
