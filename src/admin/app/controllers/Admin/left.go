// Package: Admin
// File: left.go
// Created by mint
// Useage: 管理员左侧快捷菜单
// DATE: 14-6-28 20:49
package Admin

import (
	"admin/app/models"
	"github.com/revel/revel"
	"admin/utils"
	"admin/app/controllers"
)

//左侧导航菜单
func (c *Admin) Left(menu *models.Menu) revel.Result {
	pid := c.Params.Get("pid")

	if len(pid) > 0 {
		if Pid, ok := utils.ParseMenuId(pid); ok {
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
				//获取左侧导航菜单
				left_menu := menu.GetLeftMenuHtml(Pid, admin_info)

				c.Render(left_menu)
			}
		} else {
			c.Render()
		}
	} else {
		//获取左侧导航菜单
		//默认获取 首页
		if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(1, admin_info)

			c.Render(left_menu)
		} else {
			c.Render()
		}
	}
	return c.RenderTemplate("public/left.html")
}
