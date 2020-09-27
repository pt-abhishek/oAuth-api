package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/oAuth-api/src/domain/accesstoken"
	"github.com/pt-abhishek/oAuth-api/src/domain/client"
	"github.com/pt-abhishek/oAuth-api/src/http"
	"github.com/pt-abhishek/oAuth-api/src/repository/db"
)

var (
	router = gin.Default()
)

//StartApplication gets a single instance of db and service and initialisethem
func StartApplication() {

	clientRepository := db.NewClient()
	dbRepository := db.NewAT(clientRepository)

	atService := accesstoken.NewTokenService(dbRepository)
	clientService := client.NewClientService(clientRepository)

	atHandler := http.NewATHandler(atService)
	clientHandler := http.NewClientHandler(clientService)

	//HTTP handlers for testing phase
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.POST("/oauth/client", clientHandler.CreateClient)
	router.Run(":8081")
}
