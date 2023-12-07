package routes

import (
	"log"
	"os"
	"user/handlers"
	"user/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
	redisConnectionString := os.Getenv("REDIS1")
	store, err := redis.NewStore(10, "tcp", redisConnectionString, "", []byte("secret"))
	if err != nil {
		log.Panic("redis not connected", err)
	}

	// Use sessions middleware with the Redis store
	route.Use(sessions.Sessions("mysession", store))
	route.GET("/", api.Home)
	route.GET("/user/signup", api.GetSignUp)
	route.GET("/user/login", api.GetSignIn)
	route.POST("/user/signup", api.SignUp)
	route.POST("/user/login", api.Authenticate)

	user := route.Group("/news")
	{
		//custom middleware for authorized users
		user.Use(middleware.Auth())
		user.POST("/user/preference", api.SetPreference)
	}
	return route
}
