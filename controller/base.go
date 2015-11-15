package controller

import (
	"encoding/json"
)

type ResultMap struct {
	Errorno  int32       `json:"errorno"`
	ErrorMsg string      `json:"errormsg"`
	Data     interface{} `json:"data"`
}

type BaseController struct {
}

func (c *BaseController) ApiResult(errorno int32, data interface{}) string {
	resultmap := &ResultMap{Errorno: errorno, Data: data}
	jsonbyte, err := json.Marshal(resultmap)
	checkErr(err)
	//os.Stdout.Write(groupjson)
	return string(jsonbyte)
}
