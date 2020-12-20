package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Name    string `json:"Name"`
	Host    string
	Port    string
	Mysql   Mysql
	Auth    Auth
	EnvMode string
}

type Mysql struct {
	UserName string
	Secret   string
	Ip       string
	Port     string
	DbName   string
	DB       *sql.DB
}

type Auth struct {
	AccessSecret string
}

func (sqlb *Mysql) Connect() error {
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
