package test

import (
	"errors"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis-example/transaction/db"
	"testing"
	"time"
)

type TestTable struct {
	//TableName gobatis.TableName `test_table`
	CreateTime time.Time `column:"createTime"`
	Id         int       `column:"id"`
	Password   string    `column:"password"`
	Username   string    `column:"username"`
}

func Test_InsertAutoCommit(t *testing.T) {
	testTable := TestTable{}
	gobatis.RegisterModel(&testTable)

	// 注册SQL，这里的方式就是代替了XML的写法
	gobatis.RegisterSql("insert_id", "insert into test_table (id, username, password) "+
		"values (#{TestTable.id}, #{TestTable.username}, #{TestTable.password})")
	gobatis.RegisterSql("select_id", "select * from test_table where id = #{TestTable.id}")

	var resultTestTable TestTable
	db.SessionManager.NewSession().Tx(func(session *gobatis.Session) error {
		result := 0
		err := session.Insert("insert_id").Param(TestTable{Id: 100, Username: "user5", Password: "pw5"}).Result(&result)
		if err != nil {
			return err
		}
		t.Logf("result %d\n", result)
		session.Select("select_id").Param(TestTable{Id: 100}).Result(&resultTestTable)
		t.Logf("data: %v", resultTestTable)
		//commit
		return nil
	})
}

func Test_InsertTxErrorRollback(t *testing.T) {
	testV := TestTable{}
	gobatis.RegisterModel(&testV)
	gobatis.RegisterSql("insert_id", "insert into test_table (id, username, password) "+
		"values (#{TestTable.id}, #{TestTable.username}, #{TestTable.password})")
	gobatis.RegisterSql("select_id", "select * from test_table where id = #{TestTable.id}")
	var resultTestTable TestTable
	db.SessionManager.NewSession().Tx(func(session *gobatis.Session) error {
		result := 0
		err := session.Insert("insert_id").Param(TestTable{Id: 101, Username: "user", Password: "pw"}).Result(&result)
		if err != nil {
			return err
		}
		t.Logf("ret %d\n", result)
		session.Select("select_id").Param(TestTable{Id: 101}).Result(&resultTestTable)
		t.Logf("data: %v", resultTestTable)
		// rollback
		return errors.New("rollback! ")
	})
}

func Test_InsertTxPanicRollback(t *testing.T) {
	testV := TestTable{}
	gobatis.RegisterModel(&testV)
	gobatis.RegisterSql("insert_id", "insert into test_table (id, username, password) "+
		"values (#{TestTable.id}, #{TestTable.username}, #{TestTable.password})")
	gobatis.RegisterSql("select_id", "select * from test_table where id = #{TestTable.id}")
	var resultTestTable TestTable
	db.SessionManager.NewSession().Tx(func(session *gobatis.Session) error {
		result := 0
		err := session.Insert("insert_id").Param(TestTable{Id: 102, Username: "user", Password: "pw"}).Result(&result)
		if err != nil {
			return err
		}
		t.Logf("ret %d\n", result)
		session.Select("select_id").Param(TestTable{Id: 102}).Result(&resultTestTable)
		t.Logf("data: %v", resultTestTable)

		// rollback
		panic("rollback! ")

		return nil
	})
}
