// Package: Role
// File: role.go
// Created by mint
// Useage: 角色管理
// DATE: 14-7-1 21:14
package Role

import (
	"admin/app/controllers"
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"github.com/revel/revel"
)

type Role struct {
	*revel.Controller
}

//首页
func (c *Role) Index(role *models.Role) revel.Result {

	page := c.Params.Get("page")

	if len(page) > 0 {
		if Page, ok := utils.ParsePage(page); ok {

			role_list, pages := role.GetByAll(Page, 10)
			c.Render(role_list, pages)
		}
	} else {
		role_list, pages := role.GetByAll(1, 10)

		c.Render(role_list, pages)
	}

	return c.RenderTemplate("Setting/Role/Index.html")
}

//角色成员管理
func (c *Role) Member(role *models.Role) revel.Result {

	id := c.Params.Get("id")
	page := c.Params.Get("page")

	if len(id) > 0 {
		if RoleId, ok := utils.ParseRoleId(id); ok {
			admin := new(models.Admin)

			if len(page) > 0 {
				if Page, ok := utils.ParsePage(page); ok {
					admin_list, pages := admin.GetByAll(RoleId, Page, 10)
					c.Render(admin_list, pages)
				}
			} else {
				admin_list, pages := admin.GetByAll(RoleId, 1, 10)
				c.Render(admin_list, pages)
			}
		}
	} else {
		c.Render()
	}

	return c.RenderTemplate("Setting/Role/Member.html")
}

//添加角色
func (c *Role) Add(role *models.Role) revel.Result {
	if c.Request.Method == consts.C_Method_Get {
		if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
			menu := new(models.Menu)
			tree := menu.GetMenuTree("", admin_info)

			c.Render(tree)
		} else {
			c.Render()
		}

		return c.RenderTemplate("Setting/Role/Add.html")
	} else if c.Request.Method == consts.C_Method_Post {

		rolename := c.Params.Get("rolename")
		if len(rolename) > 0 {
			role.RoleName = rolename
		} else {
			c.Flash.Error("请输入角色名称!")
			c.Flash.Out["url"] = "/Role/add/"
			return c.Redirect("/Message/")
		}

		desc := c.Params.Get("desc")
		if len(desc) > 0 {
			role.Desc = desc
		} else {
			c.Flash.Error("请输入角色描述!")
			c.Flash.Out["url"] = "/Role/add/"
			return c.Redirect("/Message/")
		}

		data := c.Params.Get("data")
		if len(data) > 0 {
			role.Data = data
		} else {
			c.Flash.Error("请选择所属权限!")
			c.Flash.Out["url"] = "/Role/add/"
			return c.Redirect("/Message/")
		}

		status := c.Params.Get("status")
		if len(status) > 0 {
			if Status, ok := utils.ParseStatus(status); ok {
				role.Status = Status
			}
		} else {
			c.Flash.Error("请选择是否启用!")
			c.Flash.Out["url"] = "/Role/add/"
			return c.Redirect("/Message/")
		}

		if role.Save() {
			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

				logs := new(models.Logs)
				desc := "添加角色:" + rolename + "|^|角色管理"
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success(c.Message("operation_success"))
			c.Flash.Out["url"] = "/Role/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error(c.Message("operation_failure"))
			c.Flash.Out["url"] = "/Role/add/"
			return c.Redirect("/Message/")
		}
	} else {

	}

	return c.Render()
}

//编辑角色
func (c *Role) Edit(role *models.Role) revel.Result {

	//当前管理员信息
	admin_info, ok := controllers.GetAdminInfoBySession(c.Session)
	if !ok {
		return c.Redirect("/Login/")
	}

	//角色ID
	id := c.Params.Get("id")
	roleId := int64(0)
	if len(id) > 0 {
		roleId, _ = utils.ParseRoleId(id)
	}

	if c.Request.Method == consts.C_Method_Get {

		if len(id) > 0 {
			role_info := role.GetById(roleId)
			menu := new(models.Menu)
			tree := menu.GetMenuTree(role_info.Data, admin_info)

			c.Render(role_info, tree, roleId)
		} else {

			menu := new(models.Menu)
			tree := menu.GetMenuTree("", admin_info)

			c.Render(tree)
		}
		return c.RenderTemplate("Setting/Role/Edit.html")
	} else if c.Request.Method == consts.C_Method_Post {

		if len(id) > 0 {

			rolename := c.Params.Get("rolename")
			if len(rolename) > 0 {
				role.RoleName = rolename
			} else {
				c.Flash.Error("请输入角色名称!")
				c.Flash.Out["url"] = "/Role/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			desc := c.Params.Get("desc")
			if len(desc) > 0 {
				role.Desc = desc
			} else {
				c.Flash.Error("请输入角色描述!")
				c.Flash.Out["url"] = "/Role/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			data := c.Params.Get("data")
			if len(data) > 0 {
				role.Data = data
			} else {
				c.Flash.Error("请选择所属权限!")
				c.Flash.Out["url"] = "/Role/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			status := c.Params.Get("status")
			if len(status) > 0 {
				Status, _ := utils.ParseStatus(status)
				role.Status = Status
			} else {
				c.Flash.Error("请选择是否启用!")
				c.Flash.Out["url"] = "/Role/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if role.Edit(roleId) {

				//******************************************
				//管理员日志
				logs := new(models.Logs)
				desc := "编辑角色:" + rolename + "|^|角色管理|^|ID:" + id
				logs.Save(admin_info, c.Controller, desc)
				//*****************************************

				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/Role/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/Role/edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error(c.Message("operation_failure"))
			c.Flash.Out["url"] = "/Role/edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	} else {

	}

	return c.Render()
}

const (
	C_SetStatus_Fail string = "0"
	C_SetStatus_Succ string = "1"
)

//设置角色状态
func (c *Role) SetStatus(role *models.Role) revel.Result {
	id := c.Params.Get("id")
	status := c.Params.Get("status")

	data := make(map[string]string)

	if len(id) > 0 && len(status) > 0 {
		roleId, _ := utils.ParseRoleId(id)
		Status, _ := utils.ParseStatus(status)

		role.Status = Status

		if role.SetStatus(roleId) {

			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
				logs := new(models.Logs)
				if Status == 1 {
					desc := "角色管理|^|设置状态|^|状态:启用"
					logs.Save(admin_info, c.Controller, desc)
				} else {
					desc := "角色管理|^|设置状态|^|状态:锁定"
					logs.Save(admin_info, c.Controller, desc)
				}
			}
			//*****************************************

			data["status"] = C_SetStatus_Succ
			data["message"] = "设置成功!"
		} else {
			data["status"] = C_SetStatus_Fail
			data["message"] = "设置失败!"
		}
	} else {
		data["status"] = C_SetStatus_Fail
		data["message"] = "设置失败!"
	}

	return c.RenderJson(data)
}

const (
	C_DelRole_Fail string = "0"
	C_DelRole_Succ string = "1"
)

//删除角色
func (c *Role) Delete(role *models.Role) revel.Result {
	id := c.Params.Get("id")
	data := make(map[string]string)

	if len(id) > 0 {
		roleId, _ := utils.ParseRoleId(id)

		if role.DelByID(roleId) {
			data["status"] = C_DelRole_Succ
			data["message"] = "删除成功!"
		} else {
			data["status"] = C_DelRole_Fail
			data["message"] = "删除失败!"
		}
	} else {
		data["status"] = C_DelRole_Fail
		data["message"] = "删除失败!"
	}

	return c.RenderJson(data)
}
