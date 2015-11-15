package controller

import (
	"database/sql"
	"fmt"
	"github.com/widuu/goini"
	"martini-simple-mvc/model"
	"net/http"
	"strings"
)

type UsersController struct {
	BaseController
}

func (c *UsersController) GetInfo(res http.ResponseWriter, req *http.Request, conf *goini.Config, db *sql.DB) string {
	// res 和 req 是通过Martini注入的
	req.ParseForm()      // 解析url传递的参数,对于POST则解析响应包的主体(request body)
	res.WriteHeader(200) // HTTP 200
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo order by uid desc limit 5")
	checkErr(err)
	var userinfo *model.UserInfo
	var groupuser []*model.UserInfo
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		userinfo = &model.UserInfo{
			Uid:        int32(uid),
			Username:   username,
			Departname: department,
			Created:    created,
		}
		groupuser = append(groupuser, userinfo)
	}
	fmt.Println(groupuser)
	apiresult := c.ApiResult(0, groupuser)
	return apiresult
}
