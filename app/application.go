package app

import (
	"github.com/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

//the only point where we use http server or Framework Gin is application handler and controller. [first layers]

//creating router
var router = gin.Default()

func StartApplication() {
	MapUrls()
	logger.Info("About to start your application")
	router.Run(":8080") // listen and serve on 8080
}
