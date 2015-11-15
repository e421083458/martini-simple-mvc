package main

import (
	"github.com/go-martini/martini"
	"martini-simple-mvc/controller"
)

func route(m *martini.ClassicMartini) {
	users := &controller.UsersController{}

	m.Get("/", func() string {
		return "Hello world!"
	})
	// m.Get("/users/:user_id", func(params martini.Params) string {
	// 	return "GET users user_id=" + params["user_id"]
	// })
	// m.Post("/users/:user_id", func(params martini.Params) string {
	// 	return "Post users user_id=" + params["user_id"]
	// })
	// m.Put("/users/:user_id", func(params martini.Params) string {
	// 	return "Put users user_id=" + params["user_id"]
	// })
	// m.Delete("/users/:user_id", func(params martini.Params) string {
	// 	return "Delete users user_id=" + params["user_id"]
	// })
	m.Get("/comm/userinfo/userinfoserv/getinfo*", users.GetInfo)
}
