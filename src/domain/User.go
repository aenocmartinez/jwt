package domain

import (
	"pulzo-login-jwt/src/infraestructure/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	repository LoginRepository
	Id         int64
	Name       string
	Email      string
	Password   string
	Active     bool
	CreatedAt  string
}

func (user *User) SetRepository(repository LoginRepository) {
	user.repository = repository
}

func (user *User) Exists() bool {
	return user.Id > 0
}

func (user *User) IsActive() bool {
	return user.Active
}

func FindUserByEmail(repository LoginRepository, email string) User {
	return repository.FindUserByEmail(email)
}

func FindUserByToken(repository LoginRepository, token string) User {
	return repository.FindUserByToken(token)
}

func (user *User) GenerateToken() string {

	config, err := config.Load("config.yml")
	if err != nil {
		panic(err)
	}

	uJwt := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		Issuer:    config.Issuer,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uJwt)

	tokenSignedString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		panic(err)
	}

	user.repository.UpdateToken(user.Id, tokenSignedString)
	return tokenSignedString
}

func (user *User) InvalidateToken() {
	user.repository.UpdateToken(user.Id, "")
}

func (user *User) Create() bool {
	return user.repository.CreateUser(*user)
}
