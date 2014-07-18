// Package: Admin
// File: logout.go
// Created by mint
// Useage: 安全退出
// DATE: 14-6-27 23:39
package Admin

import (
	"admin/app/models"
	"fmt"
	"github.com/revel/revel"
	"strconv"
)

//退出登陆
func (c *Admin) Logout(admin *models.Admin) revel.Result {

	if adminId, ok := c.Session["AdminID"]; ok {

		AdminID, err := strconv.ParseInt(adminId, 10, 64)
		if err != nil {
			revel.WARN.Printf("session解析错误: %v", err)
		}

		admin_info := admin.GetById(AdminID)

		//******************************************
		//管理员日志
		logs := new(models.Logs)
		desc := "登陆用户名:" + admin_info.UserName + "|^|退出系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
		logs.Save(admin_info, c.Controller, desc)
		//*****************************************

		for k := range c.Session {
			if k != "Lang" {
				delete(c.Session, k)
			}
		}
	}

	c.Flash.Success("安全退出")
	c.Flash.Out["url"] = "/Login/"
	return c.Redirect("/Message/")
}
