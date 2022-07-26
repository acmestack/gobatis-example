package quick_start

import (
	"encoding/json"
	"fmt"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
	"github.com/acmestack/gobatis/factory"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

var sessionManager *gobatis.SessionManager

func init() {
	err := gobatis.RegisterMapperFile("./xml/test_table_mapper.xml")
	if err != nil {
		fmt.Println("parse xml is error", err.Error())
	}
	sessionManager = gobatis.NewSessionManager(connect())
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

func TestTestTable_Insert(t *testing.T) {
	testTable := &TestTable{
		CreateTime: time.Now(),
		Username:   "user1",
		Password:   "123456",
	}
	result, id, err := testTable.Insert(sessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(id)
}

func TestTestTable_Select(t *testing.T) {
	table := &TestTable{}
	tables, err := table.Select(sessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	marshal, _ := json.Marshal(tables)
	fmt.Println(string(marshal))
}

func TestTestTable_Count(t *testing.T) {
	table := &TestTable{}
	count, err := table.Count(sessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func TestTestTable_Update(t *testing.T) {
	table := &TestTable{
		Id:       1,
		Password: "654321",
	}
	result, err := table.Update(sessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestTestTable_Delete(t *testing.T) {
	table := &TestTable{
		Username: "user1",
	}
	result, err := table.Delete(sessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
