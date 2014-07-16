// Package: models
// File: init.go
// Created by mint
// Useage: 初始化model
// DATE: 14-6-25 20:05
package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/config"
	"github.com/revel/revel"
	"os"
	"strings"
	"utils/db"
	"utils/errors"
	"utils/mail"
)

//SMTP
var (
	//系统发信
	SysMailer *mail.Mailer
)

//数据库连接
var (
	//读数据
	ReaderEngine db.DBReader
	//写数据
	WriterEngine db.DBWriter
)

func init() {
	revel.OnAppStart(initDB)
}

//初始化数据库
func initDB() {

	ReaderEngine = new(db.XormEngine)
	WriterEngine = new(db.XormEngine)

	//设置系统分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	}

	config_file := (revel.BasePath + "/conf/databases.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, _ := config.ReadDefault(config_file)

	read_driver, _ := c.String("database", "db.read.driver")
	read_dbName, _ := c.String("database", "db.read.dbname")
	read_user, _ := c.String("database", "db.read.user")
	read_password, _ := c.String("database", "db.read.password")
	read_host, _ := c.String("database", "db.read.host")
	read_encoding, _ := c.String("database", "db.read.encoding")

	//数据库连接
	var err errors.GlobalWaysError
	err = ReaderEngine.Connect(read_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", read_user, read_password, read_host, read_dbName, read_encoding))
	if err.IsError() {
		revel.WARN.Printf("数据库连接错误: %v", errors.DefaultError(err))
		os.Exit(-1)
	}

	write_driver, _ := c.String("database", "db.write.driver")
	write_dbname, _ := c.String("database", "db.write.dbname")
	write_user, _ := c.String("database", "db.write.user")
	write_password, _ := c.String("database", "db.write.password")
	write_host, _ := c.String("database", "db.write.host")
	write_encoding, _ := c.String("database", "db.write.encoding")

	err = WriterEngine.Connect(write_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", write_user, write_password, write_host, write_dbname, write_encoding))
	if err.IsError() {
		revel.WARN.Printf("数据库连接错误: %v", errors.DefaultError(err))
		os.Exit(-1)
	}

}
