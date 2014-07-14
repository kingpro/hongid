// Package: Admin
// File: admin.go
// Created by mint
// Useage: 管理员管理
// DATE: 14-6-29 23:37
package Admin

import (
	"github.com/revel/revel"
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"admin/app/controllers"
)

type Admin struct {
	*revel.Controller
}

//TODO 删除按钮应该不能使用，对于当前登录管理员

//管理员首页
func (c *Admin) Index(admin *models.Admin) revel.Result {

	page := c.Params.Get("page")

	if len(page) > 0 {
		if Page, ok := utils.ParsePage(page); ok {
			admin_list, pages := admin.GetByAll(0, Page, 10)
			c.Render(admin_list, pages)
		}
	} else {
		admin_list, pages := admin.GetByAll(0, 1, 10)
		c.Render(admin_list, pages)
	}

	return c.RenderTemplate("Setting/Admin/Index.html")
}

//添加管理员
func (c *Admin) Add(admin *models.Admin) revel.Result {

	if c.Request.Method == consts.C_Method_Get {

		role := new(models.Role)
		role_list := role.GetRoleList()

		c.Render(role_list)
		return c.RenderTemplate("Setting/Admin/Add.html")
	} else if c.Request.Method == consts.C_Method_Post {

		username := c.Params.Get("username")
		if len(username) > 0 {
			admin.Username = username
		} else {
			c.Flash.Error("请输入用户名!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		if admin.HasName() {
			c.Flash.Error("用户名“" + username + "”已存在！")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		password := c.Params.Get("password")
		if len(password) > 0 {
			admin.Password = password
		} else {
			c.Flash.Error("请输入密码!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		pwdconfirm := c.Params.Get("pwdconfirm")
		if len(pwdconfirm) > 0 {
			if password != pwdconfirm {
				c.Flash.Error("两次输入密码不一致!")
				c.Flash.Out["url"] = "/Admin/add/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("请输入确认密码!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		email := c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入E-mail!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		if admin.HasEmail() {
			c.FlashParams()
			c.Flash.Error("E-mail已存在！")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		realname := c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		lang := c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		roleid := c.Params.Get("roleid")
		if len(roleid) > 0 {
			if RoleId, ok := utils.ParseRoleId(roleid); ok {
				admin.Roleid = RoleId
			}
		} else {
			c.Flash.Error("请选择所属角色!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		status := c.Params.Get("status")
		if len(status) > 0 {
			if Status, ok := utils.ParseStatus(status); ok {
				admin.Status = Status
			}
		} else {
			c.Flash.Error("请选择状态!")
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}

		if admin.Save() {

			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
				logs := new(models.Logs)
				desc := "添加管理员:" + username + "|^|管理员管理"
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success(c.Message("operation_success"))
			c.Flash.Out["url"] = "/Admin/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error(c.Message("operation_failure"))
			c.Flash.Out["url"] = "/Admin/add/"
			return c.Redirect("/Message/")
		}
	} else {

	}

	return c.Render()
}

//编辑管理员
func (c *Admin) Edit(admin *models.Admin) revel.Result {
	if c.Request.Method == consts.C_Method_Get {

		role := new(models.Role)
		role_list := role.GetRoleList()

		id := c.Params.Get("id")

		if len(id) > 0 {
			if adminId, ok := utils.ParseAdminId(id); ok {
				admin_info := admin.GetById(adminId)
				c.Render(admin_info, role_list)
		}
		} else {
			c.Render(role_list)
		}

		return c.RenderTemplate("Setting/Admin/Edit.html")
	} else if c.Request.Method == consts.C_Method_Post {

		id := c.Params.Get("id")

		if len(id) > 0 {

			if adminId, ok := utils.ParseAdminId(id); ok {
				username := c.Params.Get("username")
				if len(username) > 0 {
					admin.Username = username
				} else {
					c.Flash.Error("请输入用户名!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				password := c.Params.Get("password")
				if len(password) > 0 {
					admin.Password = password
				}

				pwdconfirm := c.Params.Get("pwdconfirm")
				if len(pwdconfirm) > 0 {
					if password != pwdconfirm {
						c.Flash.Error("两次输入密码不一致!")
						c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
						return c.Redirect("/Message/")
					}
				}

				email := c.Params.Get("email")
				if len(email) > 0 {
					admin.Email = email
				} else {
					c.Flash.Error("请输入E-mail!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				realname := c.Params.Get("realname")
				if len(realname) > 0 {
					admin.Realname = realname
				} else {
					c.Flash.Error("请输入真实姓名!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				lang := c.Params.Get("lang")
				if len(lang) > 0 {
					admin.Lang = lang
				} else {
					c.Flash.Error("请选择语言!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				roleid := c.Params.Get("roleid")
				if len(roleid) > 0 {
					if RoleId, ok := utils.ParseRoleId(roleid); ok {
						admin.Roleid = RoleId
					}
				} else {
					c.Flash.Error("请选择所属角色!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				status := c.Params.Get("status")
				if len(status) > 0 {
					if Status, ok := utils.ParseStatus(status); ok {
						admin.Status = Status
					}
				} else {
					c.Flash.Error("请选择是否启用!")
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				//编辑
				if admin.Edit(adminId) {

					//******************************************
					//管理员日志
					if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

						logs := new(models.Logs)
						desc := "编辑管理员:" + username + "|^|管理员管理"
						logs.Save(admin_info, c.Controller, desc)
					}
					//*****************************************

					c.Flash.Success(c.Message("operation_success"))
					c.Flash.Out["url"] = "/Admin/"
					return c.Redirect("/Message/")
				} else {
					c.Flash.Error(c.Message("operation_failure"))
					c.Flash.Out["url"] = "/Admin/edit/" + id + "/"
					return c.Redirect("/Message/")
				}
			}
		} else {
			c.Flash.Error(c.Message("operation_failure"))
			c.Flash.Out["url"] = "/Admin/"
			return c.Redirect("/Message/")
		}

	} else {

	}

	return c.Render()
}


const (
	C_AdminDel_Fail string = "0"
	C_AdminDel_Succ string = "1"
)

//删除管理员
func (c *Admin) Delete(admin *models.Admin) revel.Result {
	id := c.Params.Get("id")
	data := make(map[string]string)

	if len(id) > 0 {
		if adminId, ok := utils.ParseAdminId(id); ok {
			if admin.DelByID(adminId) {

				//******************************************
				//管理员日志
				if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

					logs := new(models.Logs)
					desc := "删除管理员|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				data["status"] = C_AdminDel_Succ
				data["message"] = "删除成功!"
				return c.RenderJson(data)
			} else {
				data["status"] = C_AdminDel_Fail
				data["message"] = "删除失败!"
				return c.RenderJson(data)
			}
		} else {
			data["status"] = C_AdminDel_Fail
			data["message"] = "删除失败!"
			return c.RenderJson(data)
		}
	} else {
		data["status"] = C_AdminDel_Fail
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}

	return c.Render()
}
