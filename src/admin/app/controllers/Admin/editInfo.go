// Package: Admin
// File: editInfo.go
// Created by mint
// Useage: 修改管理员信息
// DATE: 14-6-28 21:39
package Admin

import (
	"admin/app/controllers"
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"github.com/revel/revel"
)

//个人信息更新：真实姓名，邮寄地址，系统语言
func (c *Admin) EditInfo(admin *models.Admin) revel.Result {
	if c.Request.Method == consts.C_Method_Get {
		title := "个人信息--HongID后台管理系统"

		if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {
			admin_info := admin.GetById(adminId)
			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Admin/EditInfo.html")

	} else if c.Request.Method == consts.C_Method_Post {

		//当用户登陆之后再验证其他输入信息是否正确
		if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

			//ID
			admin.Id = adminId

			//真实姓名
			realname := c.Params.Get("realname")
			if len(realname) > 0 {
				admin.Realname = realname
			} else {
				c.Flash.Error("请输入真实姓名!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}

			email := c.Params.Get("email")
			if len(email) > 0 {
				admin.Email = email
			} else {
				c.Flash.Error("请输入E-mail!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}

			//验证邮件地址未被他人注册
			if admin.HasEmail() {
				c.Flash.Error("该邮件已经注册过!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}

			//系统语言
			lang := c.Params.Get("lang")
			if len(lang) > 0 {
				admin.Lang = lang
			} else {
				c.Flash.Error("请选择语言!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}

			if admin.EditInfo(adminId) {

				//******************************************
				if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

					//更新系统语言
					c.Session[consts.C_Session_Lang] = admin_info.Lang

					//管理员日志
					logs := new(models.Logs)
					desc := "个人设置|^|个人信息"
					logs.Save(admin_info, c.Controller, desc)
				}

				if LANG, ok := c.Session[consts.C_Session_Lang]; ok {
					//设置语言
					c.Request.Locale = LANG
				} else {
					//设置默认语言
					c.Request.Locale = "zh"
				}

				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error(c.Message("not_login"))
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}

	return c.Render()
}
