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

func (m *MySQL) FindUserByEmail(email string) domain.User {
	var strQuery bytes.Buffer
	strQuery.WriteString("select id, name, email, password, active, created_at from users where email = ?")

	row := m.db.Connection().QueryRow(strQuery.String(), email)

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

func (m *MySQL) UpdateToken(idUser int64, token string) bool {
	var success bool = true
	var strQuery bytes.Buffer
	strQuery.WriteString("update users set token = ? where id = ?")

	_, err := m.db.Connection().Exec(strQuery.String(), token, idUser)
	if err != nil {
		success = false
		panic(err)
	}

	return success
}

func (m *MySQL) FindUserByToken(token string) domain.User {
	var strQuery bytes.Buffer
	strQuery.WriteString("select id, name, email, password, active, created_at from users where token = ?")

	row := m.db.Connection().QueryRow(strQuery.String(), token)

	var id int64
	var name, password, email string
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

func (m *MySQL) CreateUser(user domain.User) bool {
	var success bool = true
	var strQuery bytes.Buffer
	strQuery.WriteString("INSERT INTO users(name, email, password) VALUES (?, ?, ?)")

	_, err := m.db.Connection().Exec(strQuery.String(), user.Name, user.Email, user.Password)
	if err != nil {
		success = false
		panic(err)
	}

	return success
}
