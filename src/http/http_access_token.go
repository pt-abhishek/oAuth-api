package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/oAuth-api/src/domain/accesstoken"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//AccessTokenHandler is the interface for http layer
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	// UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

//NewATHandler New instance of http handler
func NewATHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.TokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Error parsing body to JSON")
		c.JSON(restErr.Code, restErr)
		return
	}
	token, err := h.service.Create(&at)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

// func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

// }
