package usecase

import (
	"errors"
	"pulzo-login-jwt/src/dao"
	"pulzo-login-jwt/src/domain"
)

type LogoutUseCase struct{}

func NewLogoutUseCase() *LogoutUseCase {
	return &LogoutUseCase{}
}

func (logoutCase *LogoutUseCase) Execute(encodedToken string) error {

	var repository domain.LoginRepository = dao.NewMySQL()

	user := domain.FindUserByToken(repository, encodedToken)
	if !user.Exists() {
		return errors.New("el usuario no existe")
	}
	user.SetRepository(repository)

	validateToken := NewValidateTokenUseCase()
	isValid, err := validateToken.Execute(encodedToken)
	if err != nil {
		user.InvalidateToken()
		return err
	}

	if !isValid {
		return errors.New("token no válido")
	}

	user.InvalidateToken()

	return nil
}
