package helper

import (
	"database/sql"
	"github.com/widuu/goini"
)

func NewDb() (db *sql.DB) {
	//载入config文件信息
	conf := goini.SetConfig("./config/conf.ini")
	dsn := conf.GetValue("mysql", "dsn")

	ndb := &Mysqlhelper{}
	return ndb.conn(dsn)
}
