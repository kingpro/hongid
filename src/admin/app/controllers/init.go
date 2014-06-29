// Package: controllers
// File: init.go
// Created by mint
// Useage: 初始化入口文件
// DATE: 14-6-26 14:01
package controllers

import (
	"admin/app/models"
	"github.com/revel/revel"
	"path/filepath"
	"runtime"
	"strconv"
)

var BasePath, _ = filepath.Abs("")

//定义项目根目录
var ROOT_DIR string = BasePath

//定义项目上传文件目录
var UPLOAD_DIR string = BasePath + "/www/upload/"

func init() {
	revel.OnAppStart(Bootstrap)

	//检测是否登陆
	revel.InterceptFunc(CheckLogin, revel.BEFORE, revel.ALL_CONTROLLERS)
}

//系统初始化变量
func Bootstrap() {
	//多核运行
	np := runtime.NumCPU()
	if np >= 2 {
		runtime.GOMAXPROCS(np - 1)
	}

	if runtime.GOOS == "windows" {
		UPLOAD_DIR = BasePath + "\\www\\upload\\"
	} else {
		UPLOAD_DIR = BasePath + "/www/upload/"
	}

	revel.TRACE.Println("系统初始化，不知道是不是每个连接1次还是整个服务1次")
}

//检测登陆
func CheckLogin(c *revel.Controller) revel.Result {

	//登陆页面，CSS, JS, Ajax, 验证码页面 都不进行登陆验证
	if c.Name == "Admin" && c.MethodName == "Login" || c.Name == "Ajax" || c.Name == "Static" || c.Name == "Captcha" || c.Name == "Kindeditor" {

		if LANG, ok := c.Session["Lang"]; ok {
			//设置语言
			c.RenderArgs["currentLocale"] = LANG
		} else {
			//设置默认语言
			c.RenderArgs["currentLocale"] = "zh"
		}

		revel.TRACE.Println("登陆页面，CSS, JS, Ajax, 验证码页面 都不进行登陆验证")

		return nil
	} else {
		if adminId, ok := c.Session["AdminID"]; ok {
			revel.TRACE.Println("已经保存session")
			AdminId, err := strconv.ParseInt(adminId, 10, 64)
			if err != nil {
				revel.WARN.Printf("解析Session错误: %v", err)
				return c.Redirect("/Login/")
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(AdminId)
			if admin_info.Id <= 0 {
				return c.Redirect("/Login/")
			}

			//设置语言
			c.RenderArgs["currentLocale"] = admin_info.Lang
		} else {
			revel.TRACE.Println("未保存session，跳转登录页面")
			return c.Redirect("/Login/")
		}
	}

	return nil
}
