package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/oAuth-api/src/domain/accesstoken"
	"github.com/pt-abhishek/oAuth-api/src/http"
	"github.com/pt-abhishek/oAuth-api/src/repository/db"
)

var (
	router = gin.Default()
)

//StartApplication gets a single instance of db and service and initialise them
func StartApplication() {

	dbRepository := db.New()
	atService := accesstoken.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8081")
}
