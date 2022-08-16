package controller

import (
	"log"
	"pulzo-login-jwt/src/usecase"
	"pulzo-login-jwt/src/usecase/dto"
	"pulzo-login-jwt/src/view/request"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := request.LoginRequest{}
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

func Logout(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])

	logoutUseCase := usecase.NewLogoutUseCase()
	err := logoutUseCase.Execute(tokenString)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"message": "Su sesión ha finalizado"})
}

func CreateUser(c *gin.Context) {
	req := request.CreateUserRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createUserUseCase := usecase.NewCreateUserUseCase()
	userDto := dto.UserDto{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}

	err = createUserUseCase.Execute(userDto)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
