// Package: logs
// File: logs.go
// Created by mint
// Useage: 后台操作日志记录
// DATE: 14-7-2 19:25
package logs

import (
	"admin/app/controllers"
	"admin/app/models"
	"admin/utils"
	"github.com/revel/revel"
)

type Logs struct {
	*revel.Controller
}

//日志列表
func (c *Logs) Index(logs *models.Logs) revel.Result {
	page := c.Params.Get("page")
	search := c.Params.Get("search")

	var Page int64 = 1

	if len(page) > 0 {
		Page, _ = utils.ParsePage(page)
	}

	logs_list, pages, where := logs.GetByAll(search, Page, 10)

	c.Render(logs_list, where, pages)

	return c.RenderTemplate("Setting/Logs/Index.html")
}

const (
	C_DelLog_Fail = "0"
	C_DelLog_Succ = "1"
)

//清空日志
func (c *Logs) DelAll(logs *models.Logs) revel.Result {

	data := make(map[string]string)

	IsDel := logs.DelAll()

	if IsDel {
		//******************************************
		//管理员日志
		if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
			logs := new(models.Logs)
			desc := "清空日志|^|日志管理"
			logs.Save(admin_info, c.Controller, desc)
		}
		//*****************************************

		data["status"] = C_DelLog_Succ
		data["url"] = "/Message/"
		data["message"] = "清空日志完成!"
	} else {
		data["status"] = C_DelLog_Fail
		data["url"] = "/Message/"
		data["message"] = "清空日志失败!"
	}

	return c.RenderJson(data)
}
