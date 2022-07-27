package quick_start

import (
	"encoding/json"
	"fmt"
	"github.com/acmestack/gobatis-example/transaction/db"
	"testing"
	"time"
)

func TestTestTable_Insert(t *testing.T) {
	testTable := &TestTable{
		CreateTime: time.Now(),
		Username:   "user",
		Password:   "123456",
	}
	result, id, err := testTable.Insert(db.SessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(id)
}

func TestTestTable_Select(t *testing.T) {
	table := &TestTable{}
	tables, err := table.Select(db.SessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	marshal, _ := json.Marshal(tables)
	fmt.Println(string(marshal))
}

func TestTestTable_Count(t *testing.T) {
	table := &TestTable{}
	count, err := table.Count(db.SessionManager.NewSession())
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
	result, err := table.Update(db.SessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestTestTable_Delete(t *testing.T) {
	table := &TestTable{
		Username: "user1",
	}
	result, err := table.Delete(db.SessionManager.NewSession())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
