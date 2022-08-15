package database

import (
	"database/sql"
	"fmt"
	"pulzo-login-jwt/src/infraestructure/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

func (m *MySQL) Init() {
	config, err := config.Load("config.yml")
	if err != nil {
		panic(err)
	}

	strConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	m.conn, err = sql.Open("mysql", strConn)
	if err != nil {
		panic(err)
	}

	m.conn.SetMaxIdleConns(50)
	m.conn.SetMaxOpenConns(50)
	m.conn.SetConnMaxLifetime(time.Second)
}

func (m *MySQL) Conn() *sql.DB {
	return m.conn
}

func (m *MySQL) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}
