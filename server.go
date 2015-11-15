package main

import (
	"github.com/go-martini/martini"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()
	mid(m)
	route(m)
	m.RunOnAddr(":3000")
	m.Run()
}
