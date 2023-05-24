package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vinitparekh17/go-mongo/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("/post", user.Create)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
