package main

import (
	"newsAggregator/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(api *handlers.Repository) *gin.Engine {
	route := gin.Default()
	route.Use(gin.Logger())
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https//*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
	}))
	route.GET("/fetchnews", api.GetNewsHandler)
	return route
}
