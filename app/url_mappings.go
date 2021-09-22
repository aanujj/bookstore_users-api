package app

import (
	"github.com/bookstore_users-api/controllers/ping"
	"github.com/bookstore_users-api/controllers/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	//router.GET("/users/search", user.FindUser)
}
