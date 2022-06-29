package api

import (
	"log"
	"net/http"
	"time"

	"github.com/SantiagoBedoya/supermarket_accounts-api/accounts"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	VerifyToken(c *gin.Context)
}

type handler struct {
	service accounts.Service
}

func NewHandler(service accounts.Service) Handler {
	return &handler{service}
}

func (h *handler) SignUp(c *gin.Context) {
	var user accounts.AccountSignUp
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, RestError{
			StatusCode: http.StatusBadRequest,
			Message:    accounts.InvalidUserSignUpErr.Error(),
		})
		return
	}
	account, err := h.service.SignUp(&user)
	if err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(http.StatusCreated, account.ToPublicAccount())
}

func (h *handler) VerifyToken(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		log.Println(err)
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	account, err := h.service.VerifyToken(cookie)
	if err != nil {
		restError := ParseError(err)
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(http.StatusOK, account.ToPublicAccount())
}

func (h *handler) SignIn(c *gin.Context) {
	var user accounts.AccountSignIn
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, RestError{
			StatusCode: http.StatusBadRequest,
			Message:    accounts.InvalidUserSignInErr.Error(),
		})
		return
	}
	at, err := h.service.SignIn(&user)
	if err != nil {
		restErr := ParseError(err)
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.SetCookie("jwt", at, int(time.Now().Add(time.Hour*24).Unix()), "/", "/", false, true)
	c.JSON(http.StatusOK, RestError{
		StatusCode: http.StatusOK,
		Message:    "Done",
	})
}
