package app

import (
	"github.com/annazhao/bookstore_oauth_api/src/http"
	"github.com/annazhao/bookstore_oauth_api/src/repository/db"
	"github.com/annazhao/bookstore_oauth_api/src/repository/rest"
	"github.com/annazhao/bookstore_oauth_api/src/services/accesstoken"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	// dbRepository := db.NewRepository()
	// atService := accesstoken.NewService(dbRepository) // it needs to interact with database
	// atHandler := http.NewHandler(atService)
	atHandler := http.NewAccessTokenHandler(accesstoken.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")
}
