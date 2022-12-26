package jdbc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
	"testing"
)

func init() {
	mysq_db, e := sqlx.Connect("mysql", "root:1314qqai@tcp(82.157.122.2:3306)/stars?charset=utf8&autocommit=true&parseTime=true")
	if e == nil {
		log.Println("init mysql")
		SetDataSource(mysq_db)
	} else {
		log.Println(e)
	}
}
func TestQueryForObject(t *testing.T) {
	user, err := QueryForObject[*User](func(options *QueryOptions) {
		options.Sql = "select * from user where id = ?"
		options.Args = []any{16}
	}, func(options *QueryOptions) {
		options.ElementType = new(User)
	})
	if err != nil {
		t.Error(err)
	}
	if user.Id != 16 {
		t.Error("结果不对")
	}
}
func TestQueryForMap(t *testing.T) {
	userMap, err := QueryForMap(func(options *QueryOptions) {
		options.Sql = "select * from user where id = ?"
		options.Args = []any{16}
	})
	if err != nil {
		t.Error(err)
	}
	if userMap["id"].(int64) != 16 {
		t.Error("结果不对")
	}
}

func TestQueryForMapList(t *testing.T) {
	userMap, err := QueryForMapList(func(options *QueryOptions) {
		options.Sql = "select * from user where id in (?,?)"
		options.Args = []any{16, 17}
	})
	if err != nil {
		t.Error(err)
	}
	if len(userMap) == 2 && userMap[0]["id"].(int64) != 16 {
		t.Error("结果不对")
	}
}

func TestQueryForObjectList(t *testing.T) {
	userMap, err := QueryForObjectList[*[]User](func(options *QueryOptions) {
		options.Sql = "select * from user where id in (?,?)"
		options.Args = []any{16, 17}
		ss := make([]User, 0)
		options.ElementType = &ss
	})
	if err != nil {
		t.Error(err)
	}
	if len(*userMap) == 2 && (*userMap)[0].Id != 16 {
		t.Error("结果不对")
	}
}

type User struct {
	Id               int            `db:"id"`
	UserName         sql.NullString `db:"user_name"`
	Password         sql.NullString `db:"password"`
	Phone            sql.NullString `db:"phone"`
	MembersNum       int            `db:"members_num"`
	NetworkName      sql.NullString `db:"network_name"`
	SupernodeAddress sql.NullString `db:"supernode_address"`
	Status           int            `db:"status"`
}
