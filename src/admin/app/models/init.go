// Package: models
// File: init.go
// Created by mint
// Useage: 初始化model
// DATE: 14-6-25 20:05
package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/revel/config"
	"github.com/revel/revel"
	"os"
	"strings"
	"utils/mail"
)

//数据库连接
var (
	//读数据
	DB_Read *xorm.Engine
	//写数据
	DB_Write *xorm.Engine
	//数据库前缀
	Read_prefix string
	Write_prefix string
)

//SMTP
var (
	//系统发信
	SysMailer  *mail.Mailer
)

func init() {
	revel.OnAppStart(initDB)
	revel.OnAppStart(initSmtp)
}

//初始化数据库
func initDB() {

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
	read_prefix, _ := c.String("database", "db.read.prefix")

	//数据库连接
	var err error
	DB_Read, err = xorm.NewEngine(read_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", read_user, read_password, read_host, read_dbName))
	if err != nil {
		revel.WARN.Printf("DB_Read错误: %v", err)
	}
	DB_Read.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, read_prefix))
	Read_prefix = read_prefix

	write_driver, _ := c.String("database", "db.write.driver")
	write_dbname, _ := c.String("database", "db.write.dbname")
	write_user, _ := c.String("database", "db.write.user")
	write_password, _ := c.String("database", "db.write.password")
	write_host, _ := c.String("database", "db.write.host")
	write_prefix, _ := c.String("database", "db.write.prefix")

	DB_Write, err = xorm.NewEngine(write_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", write_user, write_password, write_host, write_dbname))
	if err != nil {
		revel.WARN.Printf("DB_Write错误: %v", err)
	}
	DB_Write.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, write_prefix))
	Write_prefix = write_prefix

	//缓存方式是存放到内存中，缓存struct的记录数为1000条
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//DB_Read.SetDefaultCacher(cacher)
	//DB_Write.SetDefaultCacher(cacher)

	//控制台打印SQL语句
	//DB_Read.ShowSQL = true
	//DB_Write.ShowSQL = true

	//控制台打印调试信息
	//DB_Read.ShowDebug = true
	//DB_Write.ShowDebug = true

	//控制台打印错误信息
	//DB_Read.ShowErr = true
	//DB_Write.ShowErr = true

	//控制台打印警告信息
	//DB_Read.ShowWarn = true
	//DB_Read.ShowWarn = true
}

func initSmtp() {

	//设置系统分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	}

	config_file := (revel.BasePath + "/conf/smtp.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, _ := config.ReadDefault(config_file)

	sys_server, _ := c.String("smtp", "smtp.server")
	sys_port, _ := c.Int("smtp", "smtp.port")
	sys_userName, _ := c.String("smtp", "smtp.username")
	sys_passWord, _ := c.String("smtp", "smtp.password")
	sys_host, _ := c.String("smtp", "smtp.host")

	//配置系统发信
	SysMailer = &mail.Mailer{
		Server:     sys_server,
		Port:       sys_port,
		UserName:   sys_userName,
		Password:   sys_passWord,
		Host:       sys_host,
	}

}
