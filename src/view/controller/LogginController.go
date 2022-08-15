package controller

import (
	"pulzo-login-jwt/src/usecase"
	"pulzo-login-jwt/src/view/request"

	"github.com/gin-gonic/gin"
)

func Loggin(c *gin.Context) {
	req := request.LogginRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	loginUseCase := usecase.NewLoginUseCase()
	login, err := loginUseCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"login": login})
	}
}
