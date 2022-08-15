package database

import (
	"database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	conn *sql.DB
}

func (sqlite *SQLite) Init() {
	fmt.Println("SQLite")
	var err error
	sqlite.conn, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
}

func (sqlite *SQLite) Conn() *sql.DB {
	return sqlite.conn
}

func (sqlite *SQLite) Close() {
	if sqlite.conn != nil {
		sqlite.conn.Close()
	}
}
