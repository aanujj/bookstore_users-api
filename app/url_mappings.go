package app

import (
	"github.com/bookstore_users-api/controllers/ping"
	"github.com/bookstore_users-api/controllers/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	//to modify a given record we have to use put & when we talk about updating record we update single user
	router.PUT("/users/:user_id", user.UpdateUser)
	//to madify partial column update only use patch
	router.PATCH("/users/:user_id", user.UpdateUser)
	//to delete a user use delete  method

	router.DELETE("/users/:user_id", user.DeleteUser)
	router.GET("internal/users/search", user.FindUserStatus)
	//router.GET("/users/search", user.FindUser)
}
