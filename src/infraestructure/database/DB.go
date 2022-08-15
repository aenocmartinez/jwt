package database

import "database/sql"

type DB interface {
	Init()
	Conn() *sql.DB
	Close()
}
