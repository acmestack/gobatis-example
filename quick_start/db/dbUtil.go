package db

import (
	"fmt"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
	"github.com/acmestack/gobatis/factory"
	_ "github.com/go-sql-driver/mysql"
)

var SessionManager *gobatis.SessionManager

func init() {
	err := gobatis.RegisterMapperFile("./xml/test_table_mapper.xml")
	if err != nil {
		fmt.Println("parse xml is error:", err.Error())
	}
	SessionManager = gobatis.NewSessionManager(connect())
}

func connect() factory.Factory {
	return gobatis.NewFactory(
		gobatis.SetMaxConn(100),
		gobatis.SetMaxIdleConn(50),
		gobatis.SetDataSource(&datasource.MysqlDataSource{
			Host:     "localhost", // 数据库IP
			Port:     3306,        // 数据库端口
			DBName:   "test",      // 数据库名
			Username: "root",      // 数据库用户名
			Password: "123456",    // 数据库密码
			Charset:  "utf8",      // 编码格式
		}))
}
