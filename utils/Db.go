package utils

import (
	"database/sql"
	"fmt"
)
import _"github.com/go-sql-driver/mysql"

var Db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@/deanswjtu")
	if err != nil {
		fmt.Println("数据库连接错误:",err)
		return
	}
	err = db.Ping()
	if err == nil {
		fmt.Println("数据库连接成功")
		Db = db
		return
	}
	fmt.Println("数据库连接失败")
}
