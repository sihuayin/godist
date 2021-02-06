package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sihuayin/godist/pkg/setting"
)

var db *gorm.DB

func initData() {
	var (
		err                             error
		port, user, password, host, pre string
	)

	port = setting.DatabasePort
	user = setting.DatabaseUser
	password = setting.DatabasePWD
	host = setting.DatabaseHost
	pre = setting.DatabaseDBPRE
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
	))

	if err != nil {
		log.Fatal(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return pre + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
