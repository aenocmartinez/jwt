package usecase

import (
	"fmt"
	"pulzo-login-jwt/src/dao"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/config"

	"github.com/dgrijalva/jwt-go"
)

type ValidateTokenUseCase struct{}

func NewValidateTokenUseCase() *ValidateTokenUseCase {
	return &ValidateTokenUseCase{}
}

func (useCase *ValidateTokenUseCase) Execute(encodedToken string) (bool, error) {

	var objJwt *jwt.Token
	config, err := config.Load("config.yml")
	if err != nil {
		return false, err
	}

	var repository domain.LoginRepository = dao.NewMySQL()

	user := domain.FindUserByToken(repository, encodedToken)
	if !user.Exists() {
		return false, nil
	}

	objJwt, err = jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})

	return objJwt.Valid, err
}
