package routes

import (
	"user/handlers"
	"user/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	key := "Secret"
	store := cookie.NewStore([]byte(key))
	route.Use(sessions.Sessions("authsession", store))
	route.Use(sessions.Sessions("mysession", store))
	route.GET("/")
	route.POST("/user/signup", api.SignUp)
	route.POST("/user/login", api.Authenticate)

	user := route.Group("/news")
	{
		//custom middleware for authorized users
		user.Use(middleware.Auth())
	}

	return route
}
