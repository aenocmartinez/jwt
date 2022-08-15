package dao

import (
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/database"
)

type LoginMySQL struct {
	db *database.ConnectionPool
}

func NewLoginMySQL() *LoginMySQL {
	return &LoginMySQL{
		db: database.Instance("mysql"),
	}
}

func (lm *LoginMySQL) FindUserByEmail(email string) domain.User {
	lm.db.Connection()
	return domain.User{}
}
