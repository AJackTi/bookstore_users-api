package app

import (
	"github.com/AjackTi/bookstore_users-api/controllers/ping"
	"github.com/AjackTi/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/users/search", users.Search)
}
