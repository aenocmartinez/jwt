package database

import (
	"database/sql"
	"fmt"
	"pulzo-login-jwt/src/infraestructure/config"
	"time"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	conn *sql.DB
}

func (p *PostgreSQL) Init() {
	config, err := config.Load("config.yml")
	if err != nil {
		panic(err)
	}

	strConn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Name)

	p.conn, err = sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}

	p.conn.SetMaxIdleConns(50)
	p.conn.SetMaxOpenConns(50)
	p.conn.SetConnMaxLifetime(time.Second)
}

func (p *PostgreSQL) Conn() *sql.DB {
	return p.conn
}

func (p *PostgreSQL) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}
