package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sihuayin/godist/pkg/setting"
)

var globalDB *gorm.DB
var (
	port, user, password, host, dbName, pre string
)

func init() {
	port = setting.DatabasePort
	user = setting.DatabaseUser
	password = setting.DatabasePWD
	host = setting.DatabaseHost
	dbName = setting.DatabaseDBName
	pre = setting.DatabaseDBPRE
}

func Syncdb() {
	log.Println("数据库初始化开始")
	err := createdb()
	if err != nil {
		log.Println("数据库创建错误:", err)
		return
	}

	Connect()
	autoMigrate()
	insertUser()
	fmt.Println("数据添加完成")

}

//数据库连接
func Connect() {

	maxIdleConn := setting.DatabaseMaxIdleConn
	maxOpenConn := setting.DatabaseMaxOpenConn
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, dbName) + "&loc=Asia%2FShanghai"
	// 	//utils.Display("dbLink", dbLink)
	db, err := gorm.Open("mysql", dbLink)

	if err != nil {
		log.Fatal(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return pre + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
	globalDB = db
}

//创建数据库
func createdb() error {
	var dsn string
	var sqlstring string

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接错误:", err)
		os.Exit(2)
		//panic(err.Error())
		return err
	}
	sqlstring = fmt.Sprintf(" CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", dbName)
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
		db.Close()
		return err
	} else {
		db.Close()
		log.Println("数据库" + dbName + "创建成功")
		return nil
	}
}

// 自动迁移
func autoMigrate() {
	globalDB.AutoMigrate(&User{})
}

// 初始化用户
func insertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.ID = 1
	u.Username = "admin"
	u.IsEmailVerified = 1
	u.AuthKey = "cJIrTa_b2Hnjn6BZkrL8PJkYto2Ael3O"
	u.PasswordHash = "$2y$13$8q0MfKpnghuqCL.3FAAjiOkA8kBFNCW.ECUlqWp1zTpMHs9e5xn6u"
	u.EmailConfirmationToken = "UpToOIawm1L8GjN6pLO4r-1oj20nLT5f_1443280741"
	u.Email = "chuanzegao@163.com"
	u.Avatar = "default.jpg"
	u.Role = 1
	u.Status = 10
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Realname = "管理员"
	globalDB.Create(u)
	fmt.Println("insert user end")
}
