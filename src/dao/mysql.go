package dao

import (
	"bytes"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/database"
)

type MySQL struct {
	db *database.ConnectionPool
}

func NewMySQL() *MySQL {
	return &MySQL{
		db: database.Instance("mysql"),
	}
}

func (lm *MySQL) FindUserByEmail(email string) domain.User {
	var strQuery bytes.Buffer
	strQuery.WriteString("select id, name, email, password, active from users where email = ?")

	row := lm.db.Connection().QueryRow(strQuery.String(), email)

	var id int64
	var name, password string
	var active bool

	row.Scan(&id, &name, &email, &password, &active)

	return domain.User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Active:   active,
	}
}
