package main

import (
	"net/http"
	"pulzo-login-jwt/src/view/controller"

	"github.com/gin-gonic/gin"
)

func validateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !controller.IsValidToken(c) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Token no válido"})

		}
		c.Next()
	}
}

func main() {

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.POST("login", controller.Login)

	apiRoutes := router.Group("/pulzo-api-auth", validateTokenMiddleware())
	{
		apiRoutes.POST("/user/create", controller.CreateUser)
		apiRoutes.GET("/logout", controller.Logout)
	}

	router.Run(":8081")

}
