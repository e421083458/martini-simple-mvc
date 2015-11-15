package model

import ()

type Userinfoserv struct {
	Info string
}

func (userinfoserv *Userinfoserv) GetInfo() string {
	userinfoserv.Info = "getinfo"
	return userinfoserv.Info
}
