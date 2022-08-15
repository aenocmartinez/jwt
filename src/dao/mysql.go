package dao

import (
	"bytes"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/database"
	"time"
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
	var createdAt *time.Time

	row.Scan(&id, &name, &email, &password, &active, &createdAt)

	return domain.User{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		Active:    active,
		CreatedAt: createdAt.Format("2006-01-02"),
	}
}
