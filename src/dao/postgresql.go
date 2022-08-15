package dao

import (
	"bytes"
	"pulzo-login-jwt/src/domain"
	"pulzo-login-jwt/src/infraestructure/database"
)

type PostgreSQL struct {
	db *database.ConnectionPool
}

func NewPostgres() *PostgreSQL {
	return &PostgreSQL{
		db: database.Instance("postgres"),
	}
}

func (p *PostgreSQL) FindUserByEmail(email string) domain.User {
	var user domain.User
	var strQuery bytes.Buffer

	strQuery.WriteString("select id, name, email, password, active from users where email = $1")

	row := p.db.Connection().QueryRow(strQuery.String(), email)

	var id int64
	var name, password string
	var active bool

	row.Scan(&id, &name, &email, &password, &active)

	user = domain.User{
		Id:       id,
		Name:     name,
		Password: password,
		Active:   active,
	}

	return user
}
