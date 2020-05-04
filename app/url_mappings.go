package app

import (
	"github.com/claudiocleberson/bookstore_users-api/controllers/ping"
	"github.com/claudiocleberson/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.SearchUser)
	router.POST("users/login", users.Login)
	//router.GET("/users/search", users.SearchUser)
}
