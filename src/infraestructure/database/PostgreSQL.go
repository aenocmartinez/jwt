package database

import "database/sql"

type PostgreSQL struct {
	conn *sql.DB
}

func (p *PostgreSQL) Init() {

}

func (p *PostgreSQL) Conn() *sql.DB {
	return p.conn
}

func (p *PostgreSQL) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}
