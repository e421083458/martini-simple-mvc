package main

import (
	"github.com/go-martini/martini"
	"github.com/widuu/goini"
	"log"
	"martini-simple-mvc/helper"
)

func mid(m *martini.ClassicMartini) {
	//载入config文件信息
	conf := goini.SetConfig("./config/conf.ini")
	//port := conf.GetValue("route", "port")
	//m.RunOnAddr(":" + port)
	db := helper.NewDb()
	m.Map(conf)
	m.Map(db)

	//过滤器handle
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")
		c.Next()
		log.Println("after a request")
	})
}
