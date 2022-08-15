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
	strQuery.WriteString("select id, name, email, password, active, created_at from users where email = ?")

	row := lm.db.Connection().QueryRow(strQuery.String(), email)

	var id int64
	var name, password string
	var active bool
	var createdAt string

	row.Scan(&id, &name, &email, &password, &active, &createdAt)

	return domain.User{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		Active:    active,
		CreatedAt: createdAt,
	}
}

func (lm *MySQL) UpdateToken(idUser int64, token string) bool {
	var success bool = true
	var strQuery bytes.Buffer
	strQuery.WriteString("update users set token = ? where id = ?")

	_, err := lm.db.Connection().Exec(strQuery.String(), token, idUser)
	if err != nil {
		success = false
		panic(err)
	}

	return success
}
