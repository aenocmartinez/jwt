package database

import (
	"database/sql"
)

type ConnectionPool struct {
	db DB
}

var connectionPool *ConnectionPool

func Instance(driver string) *ConnectionPool {
	if connectionPool == nil {
		connectionPool = &ConnectionPool{
			db: factoryDB(driver),
		}
	}
	return connectionPool
}

func (cp *ConnectionPool) Connection() *sql.DB {
	return cp.db.Conn()
}

func (cp *ConnectionPool) Close() {
	if cp.db.Conn() != nil {
		cp.db.Close()
	}
}

func factoryDB(driver string) DB {
	var db DB
	db = &MySQL{}
	if driver == "postgres" {
		db = &PostgreSQL{}
	}

	db.Init()

	return db
}
