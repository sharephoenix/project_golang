package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type LYSMysql struct {
	UserName string
	Secret   string
	Ip       string
	Port     string
	DbName   string
	DB       *sql.DB
}

func (sqlb *LYSMysql) Connect() error {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		sqlb.UserName,
		sqlb.Secret,
		sqlb.Ip,
		sqlb.Port,
		sqlb.DbName)
	db, err := sql.Open("mysql", sourceName)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
		return err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	pinge := db.Ping()
	if pinge != nil {
		fmt.Println(pinge.Error())
		return err
	} else {
		fmt.Println("success!!!")
	}
	sqlb.DB = db
	return nil
}

func main() {
	mysql := LYSMysql{
		"eric",
		"qwer1234",
		"localhost",
		"3306",
		"users",
		nil,
	}
	err := mysql.Connect()
	if err != nil {
		fmt.Println("[Error]:", err.Error())
		return
	}
	defer mysql.DB.Close()
	rows, err := mysql.DB.Query("select * from users.user")
	if err != nil {
		fmt.Println("[Error-rows]:", err.Error())
		return
	}
	//defer rows.Close() //如果后面代码没有循环调用rows.Next()，就需要手动在这里释放一下，不然会一直占用缓存
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows.Columns())
	}
}
