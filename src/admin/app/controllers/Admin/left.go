// Package: Admin
// File: left.go
// Created by mint
// Useage: 管理员左侧快捷菜单
// DATE: 14-6-28 20:49
package Admin

import (
	"admin/app/models"
	"github.com/revel/revel"
	"strconv"
)

//左侧导航菜单
func (c *Admin) Left(menu *models.Menu) revel.Result {

	title := "左侧导航--HongID后台管理系统"
	pid := c.Params.Get("pid")

	if len(pid) > 0 {
		Pid, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			revel.WARN.Printf("session解析错误: %v", err)
		}

		if adminId, ok := c.Session["AdminID"]; ok {

			AdminID, err := strconv.ParseInt(adminId, 10, 64)
			if err != nil {
				revel.WARN.Printf("session解析错误: %v", err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(AdminID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(Pid, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	} else {
		//获取左侧导航菜单
		//默认获取 首页
		if adminId, ok := c.Session["AdminID"]; ok {

			AdminID, err := strconv.ParseInt(adminId, 10, 64)
			if err != nil {
				revel.WARN.Printf("session解析错误: %v", err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(AdminID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(1, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	}
	return c.RenderTemplate("public/left.html")
}
