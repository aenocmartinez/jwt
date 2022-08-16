package usecase

import (
	"errors"
	"pulzo-login-jwt/src/dao"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/bcrypt"
	"pulzo-login-jwt/src/usecase/dto"
)

type CreateUserUseCase struct{}

func NewCreateUserUseCase() *CreateUserUseCase {
	return &CreateUserUseCase{}
}

func (useCase *CreateUserUseCase) Execute(userDto dto.UserDto) error {

	var repository domain.LoginRepository = dao.NewMySQL()

	user := domain.FindUserByEmail(repository, userDto.Email)
	if user.Exists() {
		return errors.New("el usuario ya existe")
	}

	user.SetRepository(repository)

	user.Name = userDto.Name
	user.Email = userDto.Email
	user.Password = bcrypt.HashAndSalt([]byte(userDto.Password))

	_, err := user.Create()
	if err != nil {
		return err
	}

	return nil
}
