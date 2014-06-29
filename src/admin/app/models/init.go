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
)

//数据库连接

//读数据
var DB_Read *xorm.Engine

//写数据
var DB_Write *xorm.Engine

func init() {
	revel.OnAppStart(InitDB)
}

//初始化数据库
func InitDB() {

	revel.TRACE.Println("初始化DB")

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
