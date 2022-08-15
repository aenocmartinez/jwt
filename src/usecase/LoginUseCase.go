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

func (useCase *LoginUseCase) Execute(email, password string) (domain.LoginDTO, error) {

	var loginRepository domain.LoginRepository = dao.NewMySQL()
	var loginDto domain.LoginDTO

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

	loginDto.Name = user.Name
	loginDto.Email = user.Email
	loginDto.Active = user.Active
	loginDto.CreatedAt = user.CreatedAt

	return loginDto, nil
}
