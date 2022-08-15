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

func (useCase *LoginUseCase) Execute(email, password string) (LoginDTO, error) {

	var loginRepository domain.LoginRepository = dao.NewMySQL()
	var loginDto LoginDTO

	user := domain.FindUserByEmail(loginRepository, email)

	if !user.Exists() {
		return loginDto, errors.New("el usuario no existe")
	}

	if !user.IsActive() {
		return loginDto, errors.New("el usuario está inactivo")
	}

	if user.Password != password {
		return loginDto, errors.New("contraseña errada")
	}

	user.SetRepository(loginRepository)
	loginDto.Name = user.Name
	loginDto.Email = user.Email
	loginDto.Active = user.Active
	loginDto.Token = user.GenerateToken()

	return loginDto, nil
}
