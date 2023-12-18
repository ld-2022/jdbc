package jdbc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"testing"
)

var (
	DEV_MYSQL_DB = "root:1314qqai@tcp(192.168.188.48:33060)/stars?charset=utf8&autocommit=true&parseTime=true"
)

func init() {
	mysq_db, e := sqlx.Connect("mysql", DEV_MYSQL_DB)
	if e == nil {
		SetDataSource(mysq_db)
	} else {
		panic(e)
	}
}
func TestQueryForMap(t *testing.T) {
	forMap, err := QueryForMap("select * from node n where n.id = 2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(forMap)
}

func TestQueryForJSON(t *testing.T) {
	forMap, err := QueryForJSON("select * from node n where n.id = 2")
	if err != nil {
		t.Error(err)
		return
	}
	if forMap.GetInt64Value("id") != 2 {
		t.Error("id != 2")
		return
	}
	if forMap.GetString("name") != "公司测试服务器" {
		t.Error("name != 公司测试服务器")
		return
	}
	if forMap.GetString("supernode_port") != "3344" {
		t.Error("supernode_port != 3344")
		return
	}
	t.Log(forMap)
}
