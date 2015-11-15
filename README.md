# martini-simple-mvc
martini框架的简单封装

##Getting Started

安装依赖包
~~~
go get github.com/go-martini/martini
go get github.com/widuu/goini
go get github.com/go-sql-driver/mysql
~~~

然后创建数据库go_test,然后执行以下sql.
~~~
CREATE TABLE `userdetail` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `intro` varchar(2000) DEFAULT NULL,
  `profile` varchar(2000) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `userinfo` (
  `uid` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `departname` varchar(64) DEFAULT NULL,
  `created` date DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=4006 DEFAULT CHARSET=utf8;
~~~

配置config/conf.ini
~~~
dsn = root:123456@tcp(10.10.38.149:3306)/go_test
~~~

编译并且执行
~~~
go build -o server
./server
~~~

现在使用测试访问地址：
http://127.0.0.1:3000/comm/userinfo/userinfoserv/getinfo?uid=1,2&dssdadad=123232

