package main

import (
	"pulzo-login-jwt/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/pulzo-api-auth/login", controller.Login)
	router.GET("/pulzo-api-auth/logout", controller.Logout)

	router.Run(":8081")

}
