package usecase

import (
	"errors"
	"pulzo-login-jwt/src/dao"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/usecase/dto"
)

type LoginUseCase struct {
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (useCase *LoginUseCase) Execute(email, password string) (dto.LoginDTO, error) {

	var loginRepository domain.LoginRepository = dao.NewMySQL()
	var loginDto dto.LoginDTO

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

	return loginDto, nil
}
