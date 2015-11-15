package helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

type Mysqlhelper struct {
	dsn string
}

func (m *Mysqlhelper) conn(dsn string) (db *sql.DB) { // 注意
	m.dsn = dsn
	ndb, err := sql.Open("mysql", dsn)
	checkErr(err)
	return ndb
}
