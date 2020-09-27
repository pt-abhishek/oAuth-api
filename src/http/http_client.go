package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/oAuth-api/src/domain/client"
)

//ClientHandler creates clients
type ClientHandler interface {
	CreateClient(c *gin.Context)
}

type clientHandler struct {
	service client.Service
}

//NewClientHandler creates new http client Handler
func NewClientHandler(serv client.Service) ClientHandler {
	return &clientHandler{
		service: serv,
	}
}

func (h *clientHandler) CreateClient(c *gin.Context) {
	client, err := h.service.Create()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, client)

}
