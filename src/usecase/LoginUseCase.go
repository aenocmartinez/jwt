package usecase

import (
	"errors"
	"pulzo-login-jwt/src/dao"
	"pulzo-login-jwt/src/domain"
)

type LoginUseCase struct {
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (useCase *LoginUseCase) Execute(email, password string) (domain.Login, error) {

	var loginRepository domain.LoginRepository = dao.NewMySQL()
	var login domain.Login

	user := domain.FindUserByEmail(loginRepository, email)

	if !user.Exists() {
		return login, errors.New("el usuario no existe")
	}

	if !user.IsActive() {
		return login, errors.New("el usuario está inactivo")
	}

	if user.Password != password {
		return login, errors.New("contraseña errada")
	}

	user.SetRepository(loginRepository)
	login.Name = user.Name
	login.Email = user.Email
	login.Active = user.Active
	// login.CreatedAt = user.CreatedAt
	login.Token = user.GenerateToken()

	return login, nil
}
