package domain

import (
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

const APP_SECRET_KEY string = "y9T08aEM%H4d"

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

func (user *User) GenerateToken() string {
	uJwt := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		Issuer:    "pulzo.com",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uJwt)

	tokenSignedString, err := token.SignedString([]byte(APP_SECRET_KEY))
	if err != nil {
		panic(err)
	}

	user.repository.UpdateToken(user.Id, tokenSignedString)
	return tokenSignedString
}
