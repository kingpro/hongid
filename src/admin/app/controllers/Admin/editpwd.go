// Package: Admin
// File: editpwd.go
// Created by mint
// Useage: 修改管理员密码
// DATE: 14-6-29 14:26
package Admin

import (
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"github.com/revel/revel"
	"utils/security"
)

//修改密码
func (c *Admin) EditPwd(admin *models.Admin) revel.Result {
	if c.Request.Method == consts.C_Method_Get {

		if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

			admin_info := admin.GetById(adminId)

			c.Render(admin_info)
		} else {
			c.Render()
		}

		return c.RenderTemplate("Admin/EditPwd.html")
	} else {

		if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

			admin_info := admin.GetById(adminId)

			old_password := c.Params.Get("old_password")
			if len(old_password) > 0 {
				if !security.CompareHashAndPassword(admin_info.Password, old_password) {
					c.Flash.Error("旧密码不正确!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				}
			} else {
				return c.Redirect("/EditPwd/")
			}

			new_password := c.Params.Get("new_password")
			if len(new_password) <= 0 {
				c.Flash.Error("新密码不能为空!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			new_pwdconfirm := c.Params.Get("new_pwdconfirm")
			if len(new_pwdconfirm) > 0 {
				if new_pwdconfirm != new_password {
					c.Flash.Error("两次输入密码入不一致!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				} else {
					admin.Password = new_pwdconfirm
				}
			} else {
				c.Flash.Error("请重复输入新密码!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			if admin.EditPwd(adminId) {

				//******************************************
				//管理员日志
				logs := new(models.Logs)
				desc := "个人设置|^|修改密码"
				logs.Save(admin_info, c.Controller, desc)
				//*****************************************

				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error(c.Message("not_login"))
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}
